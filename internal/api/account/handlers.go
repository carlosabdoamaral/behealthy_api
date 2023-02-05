package account

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/carlosabdoamaral/behealthy_api/common"
	"github.com/carlosabdoamaral/behealthy_api/internal/responses"
	pb "github.com/carlosabdoamaral/behealthy_api/protodefs/gen/proto"
	"github.com/gin-gonic/gin"
)

func HandleUpdatePassword(ctx *gin.Context) {
	bodyBytes, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		common.LogError(err.Error())
		ctx.IndentedJSON(http.StatusConflict, err.Error())
		return
	}

	body := &responses.UpdatePasswordRequest{}
	err = json.Unmarshal(bodyBytes, body)
	if err != nil {
		common.LogError(err.Error())
		ctx.IndentedJSON(http.StatusConflict, err.Error())
		return
	}

	if body.TwoFactorCode == "" || body.NewPassword == "" {
		err = errors.New("some field is empty or isn't valid")
		common.LogError(err.Error())
		ctx.IndentedJSON(http.StatusConflict, err.Error())
		return
	}

	_, err = common.AuthServiceClient.ValidateTwoFactorCode(ctx.Request.Context(), &pb.TwoFactorCode{
		Email:         body.Email,
		TwoFactorCode: body.TwoFactorCode,
	})
	if err != nil {
		common.LogError(err.Error())
		ctx.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	status, err := common.AccountServiceClient.UpdatePassword(ctx.Request.Context(), responses.NewUpdatePasswordRequestFromJSONToProto(body))
	if err != nil {
		common.LogError(err.Error())
		ctx.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.IndentedJSON(http.StatusOK, status.GetMessage())
}

func HandleSoftDelete(ctx *gin.Context) {
	bodyBytes, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		common.LogError(err.Error())
		ctx.IndentedJSON(http.StatusConflict, err.Error())
		return
	}

	body := &responses.UpdateAccountStatusRequest{}
	err = json.Unmarshal(bodyBytes, body)
	if err != nil {
		common.LogError(err.Error())
		ctx.IndentedJSON(http.StatusConflict, err.Error())
		return
	}

	if body.TwoFactorCode == "" || body.Email == "" {
		err = errors.New("some field is empty or isn't valid")
		common.LogError(err.Error())
		ctx.IndentedJSON(http.StatusConflict, err.Error())
		return
	}

	_, err = common.AuthServiceClient.ValidateTwoFactorCode(ctx.Request.Context(), &pb.TwoFactorCode{
		Email:         body.Email,
		TwoFactorCode: body.TwoFactorCode,
	})
	if err != nil {
		common.LogError(err.Error())
		ctx.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	status, err := common.AccountServiceClient.SoftDelete(ctx.Request.Context(), responses.NewUpdateAccountStatusFromJSONToProto(body))
	if err != nil {
		common.LogError(err.Error())
		ctx.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.IndentedJSON(http.StatusOK, status.GetMessage())
}
