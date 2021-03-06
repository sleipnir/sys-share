import 'package:sys_core/pkg/shared_repositories/base_repo.dart';
import 'package:sys_share_sys_account_service/sys_share_sys_account_service.dart'
    as rpc;
import 'package:grpc/grpc_web.dart';
import 'package:shared_preferences/shared_preferences.dart';
import 'dart:convert';
import 'package:meta/meta.dart';

class AuthRepo extends BaseRepo {
  static Future<rpc.LoginResponse> loginUser(
      {@required String email, @required String password}) async {
    final req = rpc.LoginRequest()
      ..email = email
      ..password = password;

    try {
      final client = await _authClient();
      final resp = await client.login(req).then((res) {
        print(res);
        return res;
      });
      await updateTokens(
        accessToken: resp.accessToken,
        refreshToken: resp.refreshToken,
      );
      return resp;
    } catch (e) {
      print(e);
      throw e;
    }
  }

  static Future<rpc.RegisterResponse> registerAccount(
      {@required String email,
      @required String password,
      @required String passwordConfirm}) async {
    if (password != passwordConfirm) {
      rpc.RegisterResponse resp = rpc.RegisterResponse.getDefault()
        ..success = false
        ..errorReason = rpc.ErrorReason()
        ..errorReason.reason = "password mismatch";
      return resp;
    }

    try {
      final client = await _authClient();
      final request = rpc.RegisterRequest()
        ..email = email
        ..password = password
        ..passwordConfirm = passwordConfirm;
      final resp = await client.register(request);
      if (resp.hasVerifyToken()) {
        print(resp.verifyToken);
      }
      await insertTempAccountId(tempAccountId: resp.tempUserId);
      return resp;
    } catch (e) {
      throw e;
    }
  }

  static Future<rpc.ForgotPasswordResponse> forgotPassword(
      {@required String email}) async {
    final req = rpc.ForgotPasswordRequest()..email = email;
    try {
      final client = await _authClient();
      return await client.forgotPassword(req);
    } catch (e) {
      throw e;
    }
  }

  static Future<rpc.ResetPasswordResponse> resetPassword({
    @required String email,
    @required String password,
    @required String passwordConfirm,
    @required String verificationToken,
  }) async {
    final req = rpc.ResetPasswordRequest()
      ..email = email
      ..password = password
      ..passwordConfirm = passwordConfirm
      ..verifyToken = verificationToken;

    try {
      final client = await _authClient();
      return await client.resetPassword(req);
    } catch (e) {
      print(e);
      throw e;
    }
  }

  static Future<rpc.RefreshAccessTokenResponse> renewAccessToken(
      {@required String refreshToken}) async {
    final req = rpc.RefreshAccessTokenRequest()..refreshToken = refreshToken;

    try {
      final client = await _authClient();
      final resp = await client.refreshAccessToken(req);
      return resp;
    } catch (e) {
      print(e);
      throw e;
    }
  }

  static Future<void> verifyAccount(
      {@required String id, @required String verificationToken}) async {
    final req = rpc.VerifyAccountRequest()
      ..accountId = id
      ..verifyToken = verificationToken;
    try {
      final client = await _authClient();
      await client.verifyAccount(req);
    } catch (e) {
      throw e;
    }
  }

  static Future<rpc.AuthServiceClient> _authClient() async {
    return rpc.AuthServiceClient(await BaseRepo.grpcWebClientChannel());
  }
}

const _accessTokenKey = "accessToken";
const _refreshTokenKey = "refreshToken";
const _accountIdKey = "accountId";
const _tempAccountIdKey = "tempAccountId";
const _minute = 60; // 60 seconds

Future<CallOptions> getCallOptions() async {
  SharedPreferences prefs = await SharedPreferences.getInstance();
  var accToken = prefs.getString(_accessTokenKey);
  final refreshToken = prefs.getString(_refreshTokenKey);
  final currentTime = DateTime.now().toUtc().millisecondsSinceEpoch / 1000;
  if (accToken != null &&
      accToken.isNotEmpty &&
      refreshToken != null &&
      refreshToken.isNotEmpty) {
    final payload = _parseJwtPayLoad(accToken);
    final expiry = payload["exp"];
    if (expiry <= currentTime + _minute) {
      final refreshPayload = _parseJwtPayLoad(refreshToken);
      final refreshExpiry = refreshPayload["exp"];
      if (refreshExpiry <= currentTime + _minute) {
        // refreshtoken expired => log out
        await prefs.remove(_accessTokenKey);
        await prefs.remove(_refreshTokenKey);
        throw "refresh token expired, logging out";
      }
      try {
        final req = rpc.RefreshAccessTokenRequest()
          ..refreshToken = refreshToken;
        final client =
            rpc.AuthServiceClient(await BaseRepo.grpcWebClientChannel());
        client.refreshAccessToken(req).then((resp) {
          if (resp.hasAccessToken()) {
            prefs.setString(_accessTokenKey, resp.accessToken);
          }
        });
      } catch (e) {
        throw e;
      }
    }
    return CallOptions(
      metadata: {"Authorization": "Bearer " + accToken},
    );
  }
  return CallOptions();
}

Map<String, dynamic> _parseJwtPayLoad(String token) {
  final parts = token.split('.');
  if (parts.length != 3) {
    throw Exception('invalid token');
  }

  final payload = _decodeBase64(parts[1]);
  final payloadMap = json.decode(payload);
  if (payloadMap is! Map<String, dynamic>) {
    throw Exception('invalid payload');
  }

  return payloadMap;
}

String _decodeBase64(String str) {
  String output = str.replaceAll('-', '+').replaceAll('_', '/');

  switch (output.length % 4) {
    case 0:
      break;
    case 2:
      output += '==';
      break;
    case 3:
      output += '=';
      break;
    default:
      throw Exception('Illegal base64url string!"');
  }

  return utf8.decode(base64Url.decode(output));
}

Future<void> updateTokens({
  @required String accessToken,
  @required String refreshToken,
}) async {
  SharedPreferences prefs = await SharedPreferences.getInstance();
  final payload = _parseJwtPayLoad(accessToken);
  final accountId = payload["userId"];
  print(accountId);
  await prefs.setString(_accessTokenKey, accessToken);
  await prefs.setString(_refreshTokenKey, refreshToken);
  await prefs.setString(_accountIdKey, accountId);
}

Future<bool> isLoggedIn() async {
  SharedPreferences prefs = await SharedPreferences.getInstance();
  final accessToken = prefs.getString(_accessTokenKey);
  final refreshToken = prefs.getString(_refreshTokenKey);
  if (accessToken != null &&
      accessToken.isNotEmpty &&
      refreshToken != null &&
      refreshToken.isNotEmpty) {
    return true;
  }
  return false;
}

Future<String> getAccountId() async {
  SharedPreferences prefs = await SharedPreferences.getInstance();
  final accountId = prefs.getString(_accountIdKey);
  if (accountId != null && accountId.isNotEmpty) {
    return accountId;
  }
  return "";
}

Future<void> insertTempAccountId({@required String tempAccountId}) async {
  SharedPreferences prefs = await SharedPreferences.getInstance();
  prefs.setString(_tempAccountIdKey, tempAccountId);
}

Future<String> getTempAccountId() async {
  SharedPreferences prefs = await SharedPreferences.getInstance();
  final accountId = prefs.getString(_tempAccountIdKey);
  if (accountId != null && accountId.isNotEmpty) {
    return accountId;
  }
  return "";
}

Future<void> logOut() async {
  SharedPreferences prefs = await SharedPreferences.getInstance();
  prefs.remove(_accountIdKey);
  prefs.remove(_accessTokenKey);
  prefs.remove(_refreshTokenKey);
}

bool isSuperAdmin(rpc.Account account) {
  var isSuper = false;
  account.roles.forEach((role) {
    if (role.role.name == rpc.Roles.SUPERADMIN.name) {
      isSuper = true;
    }
  });
  return isSuper;
}

Map<int, rpc.UserRoles> isAdmin(rpc.Account account) {
  var mapAdminRoles = Map<int, rpc.UserRoles>();
  for (var i = 0; i < account.roles.length; i++) {
    if (account.roles[i].role == rpc.Roles.ADMIN) {
      mapAdminRoles[i] = account.roles[i];
    }
  }
  return mapAdminRoles;
}

List<String> getSubscribedOrgs(rpc.Account account) {
  var listSubscribedOrgs = List<String>();
  account.roles.forEach((role) {
    if (role.orgId.isNotEmpty) {
      listSubscribedOrgs.add(role.orgId);
    }
  });
  return listSubscribedOrgs;
}
