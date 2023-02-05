package responses

import (
	"github.com/carlosabdoamaral/behealthy_api/common"
	pb "github.com/carlosabdoamaral/behealthy_api/protodefs/gen/proto"
)

// func NewSomeModelFromJSONToProto(model *some) *pb.some {}
// func NewSomeModelFromProtoToJSON(model *pb.some) *some {}

func NewAccountDetailsFromJSONToProto(model *AccountDetails) *pb.AccountDetails {
	return &pb.AccountDetails{
		AccountId:     model.AccountId,
		FullName:      model.FullName,
		UserName:      model.UserName,
		Email:         model.Email,
		Password:      model.Password,
		BirthDate:     common.TimeToTimestamp(model.BirthDate),
		Role:          model.Role,
		CreatedAt:     common.TimeToTimestamp(model.CreatedAt),
		UpdatedAt:     common.TimeToTimestamp(model.UpdatedAt),
		SoftDeleted:   model.SoftDeleted,
		IsBlocked:     model.IsBlocked,
		TwoFactorCode: model.TwoFactorCode,
	}
}

func NewAccountDetailsFromProtoToJSON(model *pb.AccountDetails) *AccountDetails {
	return &AccountDetails{
		AccountId:     model.GetAccountId(),
		FullName:      model.GetFullName(),
		UserName:      model.GetUserName(),
		Email:         model.GetEmail(),
		Password:      model.GetPassword(),
		BirthDate:     common.TimestampToTime(model.GetBirthDate()),
		Role:          model.GetRole(),
		CreatedAt:     common.TimestampToTime(model.GetCreatedAt()),
		UpdatedAt:     common.TimestampToTime(model.GetUpdatedAt()),
		SoftDeleted:   model.GetSoftDeleted(),
		IsBlocked:     model.GetIsBlocked(),
		TwoFactorCode: model.GetTwoFactorCode(),
	}
}

func NewUpdatePasswordRequestFromJSONToProto(model *UpdatePasswordRequest) *pb.UpdatePasswordRequest {
	return &pb.UpdatePasswordRequest{
		Email:         model.Email,
		NewPassword:   model.NewPassword,
		TwoFactorCode: model.TwoFactorCode,
	}
}

func NewUpdatePasswordFromProtoToJSON(model *pb.UpdatePasswordRequest) *UpdatePasswordRequest {
	return &UpdatePasswordRequest{
		Email:         model.GetEmail(),
		NewPassword:   model.GetNewPassword(),
		TwoFactorCode: model.GetTwoFactorCode(),
	}
}
