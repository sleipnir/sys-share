package pkg

import (
	authRpc "github.com/getcouragenow/sys-share/sys-account/service/go/rpc/v2"
)

type RegisterRequest struct {
	Email           string `json:"string,omitempty"`
	Password        string `json:"password,omitempty"`
	PasswordConfirm string `json:"passwordConfirm,omitempty"`
}

func RegisterRequestFromProto(in *authRpc.RegisterRequest) *RegisterRequest {
	return &RegisterRequest{
		Email:           in.GetEmail(),
		Password:        in.GetPassword(),
		PasswordConfirm: in.GetPasswordConfirm(),
	}
}

func (r *RegisterRequest) ToProto() *authRpc.RegisterRequest {
	return &authRpc.RegisterRequest{
		Email:           r.Email,
		Password:        r.Password,
		PasswordConfirm: r.PasswordConfirm,
	}
}

type RegisterResponse struct {
	Success     bool   `json:"success,omitempty"`
	SuccessMsg  string `json:"successMsg,omitempty"`
	ErrorReason string `json:"errorReason,omitempty"`
}

func RegisterResponseFromProto(in *authRpc.RegisterResponse) *RegisterResponse {
	return &RegisterResponse{
		Success:     in.GetSuccess(),
		SuccessMsg:  in.GetSuccessMsg(),
		ErrorReason: in.GetErrorReason().GetReason(),
	}
}

func (r *RegisterResponse) ToProto() *authRpc.RegisterResponse {
	return &authRpc.RegisterResponse{
		Success:     r.Success,
		SuccessMsg:  r.SuccessMsg,
		ErrorReason: &authRpc.ErrorReason{Reason: r.ErrorReason},
	}
}

type LoginRequest struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

func (lr *LoginRequest) ToProto() *authRpc.LoginRequest {
	return &authRpc.LoginRequest{
		Email:    lr.Email,
		Password: lr.Password,
	}
}

func LoginRequestFromProto(in *authRpc.LoginRequest) *LoginRequest {
	return &LoginRequest{
		Email:    in.GetEmail(),
		Password: in.GetPassword(),
	}
}

type LoginResponse struct {
	Success      bool   `json:"success,omitempty"`
	AccessToken  string `json:"accessToken,omitempty"`
	RefreshToken string `json:"refreshToken,omitempty"`
	ErrorReason  string `json:"errorReason,omitempty"`
	LastLogin    int64  `json:"lastLogin,omitempty"`
}

func LoginResponseFromProto(in *authRpc.LoginResponse) *LoginResponse {
	return &LoginResponse{
		Success:      in.GetSuccess(),
		AccessToken:  in.GetAccessToken(),
		RefreshToken: in.GetRefreshToken(),
		ErrorReason:  in.GetErrorReason().GetReason(),
		LastLogin:    tsToUnixUTC(in.GetLastLogin()),
	}
}

func (l *LoginResponse) ToProto() *authRpc.LoginResponse {
	return &authRpc.LoginResponse{
		Success:      l.Success,
		AccessToken:  l.AccessToken,
		RefreshToken: l.RefreshToken,
		ErrorReason:  &authRpc.ErrorReason{Reason: l.ErrorReason},
		LastLogin:    unixToUtcTS(l.LastLogin),
	}
}

type ForgotPasswordRequest struct {
	Email string `json:"email,omitempty"`
}

func (fpr *ForgotPasswordRequest) ToProto() *authRpc.ForgotPasswordRequest {
	return &authRpc.ForgotPasswordRequest{Email: fpr.Email}
}

type ForgotPasswordResponse struct {
	Success                   bool   `json:"success,omitempty"`
	SuccessMsg                string `json:"successMsg,omitempty"`
	ErrorReason               string `json:"errorReason,omitempty"`
	ForgotPasswordRequestedAt int64  `json:"forgotPasswordRequestedAt,omitempty"`
}

func (fpres *ForgotPasswordResponse) ToProto() *authRpc.ForgotPasswordResponse {
	return &authRpc.ForgotPasswordResponse{
		Success:                   fpres.Success,
		SuccessMsg:                fpres.SuccessMsg,
		ErrorReason:               &authRpc.ErrorReason{Reason: fpres.ErrorReason},
		ForgotPasswordRequestedAt: unixToUtcTS(fpres.ForgotPasswordRequestedAt),
	}
}

type ResetPasswordRequest struct {
	Email           string `json:"email,omitempty"`
	Password        string `json:"password,omitempty"`
	PasswordConfirm string `json:"passwordConfirm,omitempty"`
}

func (rpr *ResetPasswordRequest) ToProto() *authRpc.ResetPasswordRequest {
	return &authRpc.ResetPasswordRequest{
		Email:           rpr.Email,
		Password:        rpr.Password,
		PasswordConfirm: rpr.PasswordConfirm,
	}
}

type ResetPasswordResponse struct {
	Success                  bool   `json:"success,omitempty"`
	SuccessMsg               string `json:"successMsg,omitempty"`
	ErrorReason              string `json:"errorReason,omitempty"`
	ResetPasswordRequestedAt int64  `json:"resetPasswordRequestedAt,omitempty"`
}

func (rps *ResetPasswordResponse) ToProto() *authRpc.ResetPasswordResponse {
	return &authRpc.ResetPasswordResponse{
		Success:                  rps.Success,
		SuccessMsg:               rps.SuccessMsg,
		ErrorReason:              &authRpc.ErrorReason{Reason: rps.ErrorReason},
		ResetPasswordRequestedAt: unixToUtcTS(rps.ResetPasswordRequestedAt),
	}
}

type RefreshAccessTokenRequest struct {
	RefreshToken string `json:"refreshToken,omitempty"`
}

func (rtr *RefreshAccessTokenRequest) ToProto() *authRpc.RefreshAccessTokenRequest {
	return &authRpc.RefreshAccessTokenRequest{RefreshToken: rtr.RefreshToken}
}

type RefreshAccessTokenResponse struct {
	AccessToken string `json:"accessToken,omitempty"`
	ErrorReason string `json:"errorReason,omitempty"`
}

func (ratr *RefreshAccessTokenResponse) ToProto() *authRpc.RefreshAccessTokenResponse {
	return &authRpc.RefreshAccessTokenResponse{
		AccessToken: ratr.AccessToken,
		ErrorReason: &authRpc.ErrorReason{Reason: ratr.ErrorReason},
	}
}
