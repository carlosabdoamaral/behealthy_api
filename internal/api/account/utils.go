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

func updateAccountStatusValidations(ctx *gin.Context) (*responses.UpdateAccountStatusRequest, error) {
	bodyBytes, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		common.LogError(err.Error())
		ctx.IndentedJSON(http.StatusConflict, err.Error())
		return nil, err
	}

	body := &responses.UpdateAccountStatusRequest{}
	err = json.Unmarshal(bodyBytes, body)
	if err != nil {
		common.LogError(err.Error())
		ctx.IndentedJSON(http.StatusConflict, err.Error())
		return nil, err
	}

	if body.TwoFactorCode == "" || body.Email == "" {
		return nil, errors.New("some field is empty or isn't valid")
	}

	_, err = common.AuthServiceClient.ValidateTwoFactorCode(ctx.Request.Context(), &pb.TwoFactorCode{
		Email:         body.Email,
		TwoFactorCode: body.TwoFactorCode,
	})

	if err != nil {
		return nil, err
	}

	return body, nil
}
