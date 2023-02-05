package responses

import (
	pb "github.com/carlosabdoamaral/behealthy_api/protodefs/gen/proto"
)

// func NewSomeModelFromJSONToProto(model *some) *pb.some {}
// func NewSomeModelFromProtoToJSON(model *pb.some) *some {}

func NewSignInRequestFromJSONToProto(model *SignInRequest) *pb.SignInRequest {
	return &pb.SignInRequest{
		Email:    model.Email,
		Password: model.Password,
	}
}

func NewSignInRequestFromProtoToJSON(model *pb.SignInRequest) *SignInRequest {
	return &SignInRequest{
		Email:    model.GetEmail(),
		Password: model.GetPassword(),
	}
}

func NewSignUpRequestFromJSONToProto(model *SignUpRequest) *pb.SignUpRequest {
	return &pb.SignUpRequest{
		FullName: model.FullName,
		UserName: model.UserName,
		Email:    model.Email,
		Password: model.Password,
		Age:      model.Age,
		Role:     model.Role,
	}
}

func NewSignUpRequestFromProtoToJSON(model *pb.SignUpRequest) *SignUpRequest {
	return &SignUpRequest{
		FullName: model.GetFullName(),
		UserName: model.GetUserName(),
		Email:    model.GetEmail(),
		Password: model.GetPassword(),
		Age:      model.GetAge(),
		Role:     model.GetRole(),
	}
}

func NewRefreshTwoFactorCodeRequestFromJSONToProto(model *RefreshTwoFactorCodeRequest) *pb.RefreshTwoFactorCodeRequest {
	return &pb.RefreshTwoFactorCodeRequest{
		Email: model.Email,
	}
}

func NewRecoverPasswordFromProtoToJSON(model *pb.UpdatePasswordRequest) *UpdatePasswordRequest {
	return &UpdatePasswordRequest{
		NewPassword:   model.GetNewPassword(),
		TwoFactorCode: model.GetTwoFactorCode(),
	}
}
