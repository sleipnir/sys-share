// Code generated by protoc-gen-cobra. DO NOT EDIT.

package v2

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	client "github.com/gutterbacon/protoc-gen-cobra/client"
	flag "github.com/gutterbacon/protoc-gen-cobra/flag"
	iocodec "github.com/gutterbacon/protoc-gen-cobra/iocodec"
	cobra "github.com/spf13/cobra"
	pflag "github.com/spf13/pflag"
	grpc "google.golang.org/grpc"
	proto "google.golang.org/protobuf/proto"
	strconv "strconv"
)

func AccountServiceClientCommand(options ...client.Option) *cobra.Command {
	cfg := client.NewConfig(options...)
	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("AccountService"),
		Short: "AccountService service client",
		Long:  "",
	}
	cfg.BindFlags(cmd.PersistentFlags())
	cmd.AddCommand(
		_AccountServiceNewAccountCommand(cfg),
		_AccountServiceGetAccountCommand(cfg),
		_AccountServiceListAccountsCommand(cfg),
		_AccountServiceSearchAccountsCommand(cfg),
		_AccountServiceAssignAccountToRoleCommand(cfg),
		_AccountServiceUpdateAccountCommand(cfg),
		_AccountServiceDisableAccountCommand(cfg),
	)
	return cmd
}

func _AccountServiceNewAccountCommand(cfg *client.Config) *cobra.Command {
	req := &Account{
		Role: &UserRoles{
			Project: &Project{},
			Org:     &Org{},
		},
		CreatedAt: &timestamp.Timestamp{},
		UpdatedAt: &timestamp.Timestamp{},
		LastLogin: &timestamp.Timestamp{},
		Fields:    &UserDefinedFields{},
	}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("NewAccount"),
		Short: "NewAccount RPC client",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "AccountService"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "AccountService", "NewAccount"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewAccountServiceClient(cc)
				v := &Account{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.NewAccount(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.Id, cfg.FlagNamer("Id"), "", "")
	cmd.PersistentFlags().StringVar(&req.Email, cfg.FlagNamer("Email"), "", "")
	cmd.PersistentFlags().StringVar(&req.Password, cfg.FlagNamer("Password"), "", "")
	_RolesVar(cmd.PersistentFlags(), &req.Role.Role, cfg.FlagNamer("Role Role"), "")
	cmd.PersistentFlags().StringVar(&req.Role.Project.Id, cfg.FlagNamer("Role Project Id"), "", "")
	cmd.PersistentFlags().StringVar(&req.Role.Org.Id, cfg.FlagNamer("Role Org Id"), "", "")
	cmd.PersistentFlags().BoolVar(&req.Role.All, cfg.FlagNamer("Role All"), false, "")
	cmd.PersistentFlags().Int64Var(&req.CreatedAt.Seconds, cfg.FlagNamer("CreatedAt Seconds"), 0, "Represents seconds of UTC time since Unix epoch\n 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n 9999-12-31T23:59:59Z inclusive.")
	cmd.PersistentFlags().Int32Var(&req.CreatedAt.Nanos, cfg.FlagNamer("CreatedAt Nanos"), 0, "Non-negative fractions of a second at nanosecond resolution. Negative\n second values with fractions must still have non-negative nanos values\n that count forward in time. Must be from 0 to 999,999,999\n inclusive.")
	cmd.PersistentFlags().Int64Var(&req.UpdatedAt.Seconds, cfg.FlagNamer("UpdatedAt Seconds"), 0, "Represents seconds of UTC time since Unix epoch\n 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n 9999-12-31T23:59:59Z inclusive.")
	cmd.PersistentFlags().Int32Var(&req.UpdatedAt.Nanos, cfg.FlagNamer("UpdatedAt Nanos"), 0, "Non-negative fractions of a second at nanosecond resolution. Negative\n second values with fractions must still have non-negative nanos values\n that count forward in time. Must be from 0 to 999,999,999\n inclusive.")
	cmd.PersistentFlags().Int64Var(&req.LastLogin.Seconds, cfg.FlagNamer("LastLogin Seconds"), 0, "Represents seconds of UTC time since Unix epoch\n 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n 9999-12-31T23:59:59Z inclusive.")
	cmd.PersistentFlags().Int32Var(&req.LastLogin.Nanos, cfg.FlagNamer("LastLogin Nanos"), 0, "Non-negative fractions of a second at nanosecond resolution. Negative\n second values with fractions must still have non-negative nanos values\n that count forward in time. Must be from 0 to 999,999,999\n inclusive.")
	cmd.PersistentFlags().BoolVar(&req.Disabled, cfg.FlagNamer("Disabled"), false, "")

	return cmd
}

func _AccountServiceGetAccountCommand(cfg *client.Config) *cobra.Command {
	req := &GetAccountRequest{}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("GetAccount"),
		Short: "GetAccount RPC client",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "AccountService"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "AccountService", "GetAccount"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewAccountServiceClient(cc)
				v := &GetAccountRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.GetAccount(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.Id, cfg.FlagNamer("Id"), "", "")

	return cmd
}

func _AccountServiceListAccountsCommand(cfg *client.Config) *cobra.Command {
	req := &ListAccountsRequest{}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("ListAccounts"),
		Short: "ListAccounts RPC client",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "AccountService"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "AccountService", "ListAccounts"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewAccountServiceClient(cc)
				v := &ListAccountsRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.ListAccounts(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().Int64Var(&req.PerPageEntries, cfg.FlagNamer("PerPageEntries"), 0, "limit")
	cmd.PersistentFlags().StringVar(&req.OrderBy, cfg.FlagNamer("OrderBy"), "", "")
	cmd.PersistentFlags().StringVar(&req.CurrentPageId, cfg.FlagNamer("CurrentPageId"), "", "number 3 => optional: current_page_token is the last id of the\n (current) listed Accounts for pagination purpose (cursor).")

	return cmd
}

func _AccountServiceSearchAccountsCommand(cfg *client.Config) *cobra.Command {
	req := &SearchAccountsRequest{
		SearchParams: &ListAccountsRequest{},
	}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("SearchAccounts"),
		Short: "SearchAccounts RPC client",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "AccountService"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "AccountService", "SearchAccounts"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewAccountServiceClient(cc)
				v := &SearchAccountsRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.SearchAccounts(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.Query, cfg.FlagNamer("Query"), "", "query can be either email, UserDefinedFields fields")
	cmd.PersistentFlags().Int64Var(&req.SearchParams.PerPageEntries, cfg.FlagNamer("SearchParams PerPageEntries"), 0, "limit")
	cmd.PersistentFlags().StringVar(&req.SearchParams.OrderBy, cfg.FlagNamer("SearchParams OrderBy"), "", "")
	cmd.PersistentFlags().StringVar(&req.SearchParams.CurrentPageId, cfg.FlagNamer("SearchParams CurrentPageId"), "", "number 3 => optional: current_page_token is the last id of the\n (current) listed Accounts for pagination purpose (cursor).")

	return cmd
}

func _AccountServiceAssignAccountToRoleCommand(cfg *client.Config) *cobra.Command {
	req := &AssignAccountToRoleRequest{
		Role: &UserRoles{
			Project: &Project{},
			Org:     &Org{},
		},
	}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("AssignAccountToRole"),
		Short: "AssignAccountToRole RPC client",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "AccountService"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "AccountService", "AssignAccountToRole"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewAccountServiceClient(cc)
				v := &AssignAccountToRoleRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.AssignAccountToRole(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.AssigneeAccountId, cfg.FlagNamer("AssigneeAccountId"), "", "")
	cmd.PersistentFlags().StringVar(&req.AssignedAccountId, cfg.FlagNamer("AssignedAccountId"), "", "")
	_RolesVar(cmd.PersistentFlags(), &req.Role.Role, cfg.FlagNamer("Role Role"), "")
	cmd.PersistentFlags().StringVar(&req.Role.Project.Id, cfg.FlagNamer("Role Project Id"), "", "")
	cmd.PersistentFlags().StringVar(&req.Role.Org.Id, cfg.FlagNamer("Role Org Id"), "", "")
	cmd.PersistentFlags().BoolVar(&req.Role.All, cfg.FlagNamer("Role All"), false, "")

	return cmd
}

func _AccountServiceUpdateAccountCommand(cfg *client.Config) *cobra.Command {
	req := &Account{
		Role: &UserRoles{
			Project: &Project{},
			Org:     &Org{},
		},
		CreatedAt: &timestamp.Timestamp{},
		UpdatedAt: &timestamp.Timestamp{},
		LastLogin: &timestamp.Timestamp{},
		Fields:    &UserDefinedFields{},
	}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("UpdateAccount"),
		Short: "UpdateAccount RPC client",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "AccountService"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "AccountService", "UpdateAccount"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewAccountServiceClient(cc)
				v := &Account{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.UpdateAccount(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.Id, cfg.FlagNamer("Id"), "", "")
	cmd.PersistentFlags().StringVar(&req.Email, cfg.FlagNamer("Email"), "", "")
	cmd.PersistentFlags().StringVar(&req.Password, cfg.FlagNamer("Password"), "", "")
	_RolesVar(cmd.PersistentFlags(), &req.Role.Role, cfg.FlagNamer("Role Role"), "")
	cmd.PersistentFlags().StringVar(&req.Role.Project.Id, cfg.FlagNamer("Role Project Id"), "", "")
	cmd.PersistentFlags().StringVar(&req.Role.Org.Id, cfg.FlagNamer("Role Org Id"), "", "")
	cmd.PersistentFlags().BoolVar(&req.Role.All, cfg.FlagNamer("Role All"), false, "")
	cmd.PersistentFlags().Int64Var(&req.CreatedAt.Seconds, cfg.FlagNamer("CreatedAt Seconds"), 0, "Represents seconds of UTC time since Unix epoch\n 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n 9999-12-31T23:59:59Z inclusive.")
	cmd.PersistentFlags().Int32Var(&req.CreatedAt.Nanos, cfg.FlagNamer("CreatedAt Nanos"), 0, "Non-negative fractions of a second at nanosecond resolution. Negative\n second values with fractions must still have non-negative nanos values\n that count forward in time. Must be from 0 to 999,999,999\n inclusive.")
	cmd.PersistentFlags().Int64Var(&req.UpdatedAt.Seconds, cfg.FlagNamer("UpdatedAt Seconds"), 0, "Represents seconds of UTC time since Unix epoch\n 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n 9999-12-31T23:59:59Z inclusive.")
	cmd.PersistentFlags().Int32Var(&req.UpdatedAt.Nanos, cfg.FlagNamer("UpdatedAt Nanos"), 0, "Non-negative fractions of a second at nanosecond resolution. Negative\n second values with fractions must still have non-negative nanos values\n that count forward in time. Must be from 0 to 999,999,999\n inclusive.")
	cmd.PersistentFlags().Int64Var(&req.LastLogin.Seconds, cfg.FlagNamer("LastLogin Seconds"), 0, "Represents seconds of UTC time since Unix epoch\n 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n 9999-12-31T23:59:59Z inclusive.")
	cmd.PersistentFlags().Int32Var(&req.LastLogin.Nanos, cfg.FlagNamer("LastLogin Nanos"), 0, "Non-negative fractions of a second at nanosecond resolution. Negative\n second values with fractions must still have non-negative nanos values\n that count forward in time. Must be from 0 to 999,999,999\n inclusive.")
	cmd.PersistentFlags().BoolVar(&req.Disabled, cfg.FlagNamer("Disabled"), false, "")

	return cmd
}

func _AccountServiceDisableAccountCommand(cfg *client.Config) *cobra.Command {
	req := &DisableAccountRequest{}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("DisableAccount"),
		Short: "DisableAccount RPC client",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "AccountService"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "AccountService", "DisableAccount"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewAccountServiceClient(cc)
				v := &DisableAccountRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.DisableAccount(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.AccountId, cfg.FlagNamer("AccountId"), "", "")

	return cmd
}

type _RolesValue Roles

func _RolesVar(fs *pflag.FlagSet, p *Roles, name, usage string) {
	fs.Var((*_RolesValue)(p), name, usage)
}

func (v *_RolesValue) Set(val string) error {
	if e, err := parseRoles(val); err != nil {
		return err
	} else {
		*v = _RolesValue(e)
		return nil
	}
}

func (*_RolesValue) Type() string { return "Roles" }

func (v *_RolesValue) String() string { return (Roles)(*v).String() }

func parseRoles(s string) (Roles, error) {
	if i, ok := Roles_value[s]; ok {
		return Roles(i), nil
	} else if i, err := strconv.ParseInt(s, 0, 32); err == nil {
		return Roles(i), nil
	} else {
		return 0, err
	}
}

func AuthServiceClientCommand(options ...client.Option) *cobra.Command {
	cfg := client.NewConfig(options...)
	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("AuthService"),
		Short: "AuthService service client",
		Long:  "",
	}
	cfg.BindFlags(cmd.PersistentFlags())
	cmd.AddCommand(
		_AuthServiceRegisterCommand(cfg),
		_AuthServiceLoginCommand(cfg),
		_AuthServiceForgotPasswordCommand(cfg),
		_AuthServiceResetPasswordCommand(cfg),
		_AuthServiceRefreshAccessTokenCommand(cfg),
	)
	return cmd
}

func _AuthServiceRegisterCommand(cfg *client.Config) *cobra.Command {
	req := &RegisterRequest{}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("Register"),
		Short: "Register RPC client",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "AuthService"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "AuthService", "Register"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewAuthServiceClient(cc)
				v := &RegisterRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.Register(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.Email, cfg.FlagNamer("Email"), "", "")
	cmd.PersistentFlags().StringVar(&req.Password, cfg.FlagNamer("Password"), "", "")
	cmd.PersistentFlags().StringVar(&req.PasswordConfirm, cfg.FlagNamer("PasswordConfirm"), "", "")

	return cmd
}

func _AuthServiceLoginCommand(cfg *client.Config) *cobra.Command {
	req := &LoginRequest{}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("Login"),
		Short: "Login RPC client",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "AuthService"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "AuthService", "Login"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewAuthServiceClient(cc)
				v := &LoginRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.Login(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.Email, cfg.FlagNamer("Email"), "", "")
	cmd.PersistentFlags().StringVar(&req.Password, cfg.FlagNamer("Password"), "", "")

	return cmd
}

func _AuthServiceForgotPasswordCommand(cfg *client.Config) *cobra.Command {
	req := &ForgotPasswordRequest{}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("ForgotPassword"),
		Short: "ForgotPassword RPC client",
		Long:  "ForgotPassword, then ResetPassword if succeed",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "AuthService"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "AuthService", "ForgotPassword"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewAuthServiceClient(cc)
				v := &ForgotPasswordRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.ForgotPassword(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.Email, cfg.FlagNamer("Email"), "", "")

	return cmd
}

func _AuthServiceResetPasswordCommand(cfg *client.Config) *cobra.Command {
	req := &ResetPasswordRequest{}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("ResetPassword"),
		Short: "ResetPassword RPC client",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "AuthService"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "AuthService", "ResetPassword"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewAuthServiceClient(cc)
				v := &ResetPasswordRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.ResetPassword(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.Email, cfg.FlagNamer("Email"), "", "")
	cmd.PersistentFlags().StringVar(&req.Password, cfg.FlagNamer("Password"), "", "")
	cmd.PersistentFlags().StringVar(&req.PasswordConfirm, cfg.FlagNamer("PasswordConfirm"), "", "")

	return cmd
}

func _AuthServiceRefreshAccessTokenCommand(cfg *client.Config) *cobra.Command {
	req := &RefreshAccessTokenRequest{}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("RefreshAccessToken"),
		Short: "RefreshAccessToken RPC client",
		Long:  "Refresh Access Token endpoint",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "AuthService"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "AuthService", "RefreshAccessToken"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewAuthServiceClient(cc)
				v := &RefreshAccessTokenRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.RefreshAccessToken(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.RefreshToken, cfg.FlagNamer("RefreshToken"), "", "")

	return cmd
}
