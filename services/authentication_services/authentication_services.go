package authentication_services

import (
	"aapanavyapar_service_authentication/data_base/data_services"
	"aapanavyapar_service_authentication/data_base/helpers"
	"aapanavyapar_service_authentication/data_base/structs"
	"aapanavyapar_service_authentication/pb"
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"github.com/o1egl/paseto/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"os"
	"time"
)

type AuthenticationServer struct {
	data *data_services.DataServices
}

func NewAuthenticationServer() (*AuthenticationServer, error) {
	return &AuthenticationServer{
		data: data_services.NewDbConnection(),
	}, nil
}

func PrintClaimsOfAuthToken(token string) {
	var newJsonToken paseto.JSONToken
	var newFooter string
	err := paseto.Decrypt(token, []byte(os.Getenv("AUTH_TOKEN_SECRETE")), &newJsonToken, &newFooter)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Auth Token")
		fmt.Println("Audience", newJsonToken.Audience)
		fmt.Println("Subject : ", newJsonToken.Subject)
		fmt.Println("Expiration : ", newJsonToken.Expiration)
		fmt.Println("IssueAt : ", newJsonToken.IssuedAt)
		fmt.Println("Issuer : ", newJsonToken.Issuer)
		var val bool
		_ = newJsonToken.Get("authorized", &val)
		fmt.Println("Authorized : ", val)
		fmt.Println("Footer : ", newFooter)
	}
}

func PrintClaimsOfRefreshToken(token string) {
	var newJsonToken paseto.JSONToken
	var newFooter string
	err := paseto.Decrypt(token, []byte(os.Getenv("REFRESH_TOKEN_SECRETE")), &newJsonToken, &newFooter)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Refresh Token")
		fmt.Println("Audience", newJsonToken.Audience)
		fmt.Println("Subject : ", newJsonToken.Subject)
		fmt.Println("Expiration : ", newJsonToken.Expiration)
		fmt.Println("IssueAt : ", newJsonToken.IssuedAt)
		fmt.Println("Issuer : ", newJsonToken.Issuer)
		var val bool
		_ = newJsonToken.Get("authorized", &val)
		fmt.Println("Authorized : ", val)
		fmt.Println("Footer : ", newFooter)
	}
}

func (authenticationServer *AuthenticationServer) GetNewToken(ctx context.Context, request *pb.NewTokenRequest) (*pb.NewTokenResponse, error) {
	existingRefreshToken := request.GetRefreshToken()
	if existingRefreshToken == "" {
		return nil, status.Errorf(codes.InvalidArgument, "No Refresh Token Is Specified")
	}

	ok, token, err := authenticationServer.data.ValidateRefreshTokenAndGenerateNewAuthToken(ctx, existingRefreshToken)
	if err != nil {
		return nil, err
	}

	if ok {
		return &pb.NewTokenResponse{
			Token: token,
		}, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "Invalid Token")
}

func (authenticationServer *AuthenticationServer) Signup(ctx context.Context, request *pb.SignUpRequest) (*pb.SignUpResponse, error) {

	user, err := helpers.SanitizeAndValidate(request)
	fmt.Println("Sanitization and validation completed")

	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.Code(pb.ProblemCode_NoUserNameIsProvided):
				return &pb.SignUpResponse{
					Data: &pb.SignUpResponse_Code{
						Code: pb.ProblemCode_NoUserNameIsProvided,
					},
					Authorized: false,
				}, nil
			case codes.Code(pb.ProblemCode_NoPhoneNumberIsProvided):
				return &pb.SignUpResponse{
					Data: &pb.SignUpResponse_Code{
						Code: pb.ProblemCode_NoPhoneNumberIsProvided,
					},
					Authorized: false,
				}, nil
			case codes.Code(pb.ProblemCode_NoPasswordIsProvided):
				return &pb.SignUpResponse{
					Data: &pb.SignUpResponse_Code{
						Code: pb.ProblemCode_NoPasswordIsProvided,
					},
					Authorized: false,
				}, nil
			case codes.Code(pb.ProblemCode_InvalidPasswordLength):
				return &pb.SignUpResponse{
					Data: &pb.SignUpResponse_Code{
						Code: pb.ProblemCode_InvalidPasswordLength,
					},
					Authorized: false,
				}, nil
			case codes.Code(pb.ProblemCode_InvalidPhoneNumber):
				return &pb.SignUpResponse{
					Data: &pb.SignUpResponse_Code{
						Code: pb.ProblemCode_InvalidPhoneNumber,
					},
					Authorized: false,
				}, nil
			case codes.Code(pb.ProblemCode_InvalidPinCode):
				return &pb.SignUpResponse{
					Data: &pb.SignUpResponse_Code{
						Code: pb.ProblemCode_InvalidPinCode,
					},
					Authorized: false,
				}, nil
			case codes.Code(pb.ProblemCode_InvalidEmailAddress):
				return &pb.SignUpResponse{
					Data: &pb.SignUpResponse_Code{
						Code: pb.ProblemCode_InvalidEmailAddress,
					},
					Authorized: false,
				}, nil
			}
		}
		return nil, err
	}

	if _, err := authenticationServer.data.GetContactListDataFormCash(ctx, user.PhoneNo); err == nil {
		return &pb.SignUpResponse{
			Data: &pb.SignUpResponse_Code{
				Code: pb.ProblemCode_UserAlreadyExistWithSameContactNumber,
			},
			Authorized: false,
		}, nil

	}

	userId, err := uuid.NewRandom()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "can not generate internal userId  : %w", err)
	}
	fmt.Println("UUID Generated")

	err = authenticationServer.data.CreateTemporaryUserInCash(ctx, &structs.UserData{
		UserId:   userId.String(),
		Username: user.GetUsername(),
		Password: user.GetPassword(),
		PhoneNo:  user.GetPhoneNo(),
		Email:    user.GetEmail(),
		PinCode:  user.GetPinCode(),
	})
	if err != nil {
		return nil, err
	}

	refreshToken, authToken, err := authenticationServer.data.GenerateRefreshAndAuthTokenAndAddRefreshToCash(ctx, userId.String(), false)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Unable To Generate Refresh Token", err)
	}

	err = authenticationServer.data.GenerateAndSendOTP(ctx, userId.String(), user.GetPhoneNo(), 0, data_services.Validation5Min)
	if err != nil {
		return nil, err
	}

	return &pb.SignUpResponse{
		Data: &pb.SignUpResponse_ResponseData{
			ResponseData: &pb.ResponseData{
				Token:        authToken,
				RefreshToken: refreshToken,
			},
		},
		Authorized: false,
	}, nil
}

func (authenticationServer *AuthenticationServer) SignInWithMail(ctx context.Context, request *pb.SignInForMailBaseRequest) (*pb.SignInForMailBaseResponse, error) {
	email, err := helpers.SanitizeAndValidateEmailAddress(request.Mail)
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.Code(pb.ProblemCode_NoEmailIsProvided):
				return &pb.SignInForMailBaseResponse{
					Data: &pb.SignInForMailBaseResponse_Code{
						Code: pb.ProblemCode_NoEmailIsProvided,
					},
				}, nil
			case codes.Code(pb.ProblemCode_InvalidEmailAddress):
				return &pb.SignInForMailBaseResponse{
					Data: &pb.SignInForMailBaseResponse_Code{
						Code: pb.ProblemCode_InvalidEmailAddress,
					},
				}, nil
			}
		}
	}
	password, err := helpers.SanitizeAndValidatePassword(request.Password)
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.Code(pb.ProblemCode_NoPasswordIsProvided):
				return &pb.SignInForMailBaseResponse{
					Data: &pb.SignInForMailBaseResponse_Code{
						Code: pb.ProblemCode_NoPasswordIsProvided,
					},
				}, nil
			case codes.Code(pb.ProblemCode_InvalidPasswordLength):
				return &pb.SignInForMailBaseResponse{
					Data: &pb.SignInForMailBaseResponse_Code{
						Code: pb.ProblemCode_InvalidPasswordLength,
					},
				}, nil
			}
		}
	}
	fmt.Println("Sanitization and validation completed")

	userId, err := authenticationServer.data.SignInWithMailAndPassword(email, password)
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.Code(pb.ProblemCode_InvalidUserCredentials):
				return &pb.SignInForMailBaseResponse{
					Data: &pb.SignInForMailBaseResponse_Code{
						Code: pb.ProblemCode_InvalidUserCredentials,
					},
				}, nil

			case codes.Code(pb.ProblemCode_InvalidPassword):
				return &pb.SignInForMailBaseResponse{
					Data: &pb.SignInForMailBaseResponse_Code{
						Code: pb.ProblemCode_InvalidPassword,
					},
				}, nil

			}
		}
		return nil, status.Errorf(codes.Unknown, "Unable To Authenticate", err)
	}

	refreshToken, authToken, err := authenticationServer.data.GenerateRefreshAndAuthTokenAndAddRefreshToCash(ctx, userId, true)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Unable To Generate Refresh Token", err)
	}

	return &pb.SignInForMailBaseResponse{
		Data: &pb.SignInForMailBaseResponse_ResponseData{
			ResponseData: &pb.ResponseData{
				Token:        authToken,
				RefreshToken: refreshToken,
			},
		},
	}, nil

}

func (authenticationServer *AuthenticationServer) Logout(ctx context.Context, request *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	token, err := authenticationServer.data.ValidateToken(ctx, request.GetToken(), os.Getenv("AUTH_TOKEN_SECRETE"))
	if err != nil {
		return &pb.LogoutResponse{Status: false}, err
	}

	err = authenticationServer.data.DelDataFormCash(ctx, token.Subject)
	if err != nil {
		return &pb.LogoutResponse{Status: false}, err
	}
	return &pb.LogoutResponse{Status: true}, nil
}

func (authenticationServer *AuthenticationServer) ContactConformation(ctx context.Context, request *pb.ContactConformationRequest) (*pb.ContactConformationResponse, error) {
	token, err := authenticationServer.data.ValidateToken(ctx, request.GetToken(), os.Getenv("AUTH_TOKEN_SECRETE"))
	if err != nil {
		return nil, err
	}

	var authorized bool
	if err = token.Get("authorized", &authorized); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid Token", err)
	}
	if authorized {
		return nil, status.Errorf(codes.AlreadyExists, "Already Authorized")
	}

	cashVal, err := authenticationServer.data.GetDataFormCash(ctx, token.Audience)
	if err != nil {
		return nil, err
	}

	fmt.Println("Val : ", cashVal)
	fmt.Println("Requested OTP : ", request.GetOtp())

	var val structs.OTPCashData
	structs.UnmarshalOTPCash([]byte(cashVal), &val)

	if val.OTP == request.GetOtp() {

		data, err := authenticationServer.data.GetTemporaryUserFromCash(ctx, token.Audience)
		if err != nil {
			return nil, err
		}

		err = authenticationServer.data.CreateUser(ctx, data)
		if err != nil {
			return nil, err // If User Already Exist Then Report Inconsistency with cash and database
		}

		err = authenticationServer.data.DelDataFormCash(ctx, token.Subject)
		if err != nil {
			return nil, err
		}
		fmt.Println("Refresh From Cash Delete ", err)

		err = authenticationServer.data.DelDataFormCash(ctx, token.Audience)
		if err != nil {
			return nil, err
		}
		fmt.Println("Token From Cash Delete ", err)

		refreshTok, authTok, err := authenticationServer.data.GenerateRefreshAndAuthTokenAndAddRefreshToCash(ctx, token.Audience, true)
		if err != nil {
			return nil, err
		}

		return &pb.ContactConformationResponse{
			Token:        authTok,
			RefreshToken: refreshTok,
		}, nil

	}
	return nil, status.Errorf(codes.InvalidArgument, "Invalid OTP")
}

func (authenticationServer *AuthenticationServer) ResendOTP(ctx context.Context, request *pb.ResendOTPRequest) (*pb.ResendOTPResponse, error) {
	token, err := authenticationServer.data.ValidateToken(ctx, request.GetToken(), os.Getenv("AUTH_TOKEN_SECRETE"))
	if err != nil {
		return nil, err
	}

	var authorized bool
	if err = token.Get("authorized", &authorized); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid Token", err)
	}
	if authorized {
		return nil, status.Errorf(codes.PermissionDenied, "You are not authorized for this service")
	}

	val, err := authenticationServer.data.GetDataFormCash(ctx, token.Audience)

	var data structs.OTPCashData
	structs.UnmarshalOTPCash([]byte(val), &data)

	fmt.Println("Data Resend Times : ", data.ResendTimes)
	fmt.Println("Time Of OTP Sending : ", data.Time)
	fmt.Println("Current Time  : ", time.Now())

	// If OTPResponse_Ok then TimeToWaitForNextRequest is time after which you can get *next* otp if required
	// If OTPResponse_NotOk then TimeToWaitForNextRequest is time to wait to get otp.

	switch data.ResendTimes {
	case 0:
		err = authenticationServer.data.GenerateAndSendOTP(ctx, token.Audience, data.PhoneNo, 1, data_services.Validation5Min+time.Minute)
		if err != nil {
			return nil, err
		}

		return &pb.ResendOTPResponse{
			Response:                 pb.OTPResponse_OK,
			TimeToWaitForNextRequest: ptypes.DurationProto(data_services.Validation5Min),
		}, nil

	case 1:
		if time.Now().Sub(data.Time) >= data_services.Validation5Min {
			err = authenticationServer.data.GenerateAndSendOTP(ctx, token.Audience, data.PhoneNo, 2, data_services.Validation5Min+time.Minute)
			if err != nil {
				return nil, err
			}

			return &pb.ResendOTPResponse{
				Response:                 pb.OTPResponse_OK,
				TimeToWaitForNextRequest: ptypes.DurationProto(data_services.Validation5Min),
			}, nil
		}

		return &pb.ResendOTPResponse{
			Response:                 pb.OTPResponse_NotOk,
			TimeToWaitForNextRequest: ptypes.DurationProto(time.Now().Sub(data.Time)),
		}, nil

	case 2:
		if time.Now().Sub(data.Time) >= data_services.Validation5Min {
			err = authenticationServer.data.GenerateAndSendOTP(ctx, token.Audience, data.PhoneNo, 3, data_services.Validation5Min+time.Minute)
			if err != nil {
				return nil, err
			}

			return &pb.ResendOTPResponse{
				Response:                 pb.OTPResponse_OK,
				TimeToWaitForNextRequest: ptypes.DurationProto(data_services.Validation10Min),
			}, nil
		}

		return &pb.ResendOTPResponse{
			Response:                 pb.OTPResponse_NotOk,
			TimeToWaitForNextRequest: ptypes.DurationProto(time.Now().Sub(data.Time)),
		}, nil

	case 3:
		if time.Now().Sub(data.Time) >= data_services.Validation10Min {
			err = authenticationServer.data.GenerateAndSendOTP(ctx, token.Audience, data.PhoneNo, 4, data_services.Validation10Min+time.Minute)
			if err != nil {
				return nil, err
			}

			return &pb.ResendOTPResponse{
				Response:                 pb.OTPResponse_OK,
				TimeToWaitForNextRequest: ptypes.DurationProto(data_services.Validation5Hr),
			}, nil
		}

		return &pb.ResendOTPResponse{
			Response:                 pb.OTPResponse_NotOk,
			TimeToWaitForNextRequest: ptypes.DurationProto(time.Now().Sub(data.Time)),
		}, nil

	case 4:
		if time.Now().Sub(data.Time) >= data_services.Validation12Hr {
			err = authenticationServer.data.GenerateAndSendOTP(ctx, token.Audience, data.PhoneNo, 4, data_services.Validation12Hr+time.Minute)
			if err != nil {
				return nil, err
			}

			return &pb.ResendOTPResponse{
				Response:                 pb.OTPResponse_OK,
				TimeToWaitForNextRequest: ptypes.DurationProto(data_services.Validation12Hr),
			}, nil
		}

		return &pb.ResendOTPResponse{
			Response:                 pb.OTPResponse_NotOk,
			TimeToWaitForNextRequest: ptypes.DurationProto(time.Now().Sub(data.Time)),
		}, nil

	}
	return nil, status.Errorf(codes.Unknown, "Unable To Process Request")
}
