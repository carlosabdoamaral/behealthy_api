package auth

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/carlosabdoamaral/behealthy_api/common"
	"github.com/carlosabdoamaral/behealthy_api/internal/responses"
	"github.com/gin-gonic/gin"
)

func HandleSignIn(ctx *gin.Context) {
	bodyBytes, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		common.LogError(err.Error())
		ctx.IndentedJSON(http.StatusConflict, err.Error())
		return
	}

	body := &responses.SignInRequest{}
	err = json.Unmarshal(bodyBytes, body)
	if err != nil {
		common.LogError(err.Error())
		ctx.IndentedJSON(http.StatusConflict, err.Error())
		return
	}

	if body.Email == "" || body.Password == "" {
		err = errors.New("email and password mustn't be empty")
		common.LogError(err.Error())
		ctx.IndentedJSON(http.StatusConflict, err.Error())
		return
	}

	if !common.EmailIsValid(body.Email) {
		err = errors.New("email isn't valid")
		common.LogError(err.Error())
		ctx.IndentedJSON(http.StatusConflict, err.Error())
		return
	}

	protoMsg := responses.NewSignInRequestFromJSONToProto(body)
	protoResp, err := common.AuthServiceClient.SignIn(ctx.Request.Context(), protoMsg)
	if err != nil {
		common.LogError(err.Error())
		ctx.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	jsonResp := responses.NewAccountDetailsFromProtoToJSON(protoResp)
	ctx.IndentedJSON(http.StatusOK, jsonResp)
}

func HandleSignUp(ctx *gin.Context) {
	bodyBytes, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		common.LogError(err.Error())
		ctx.IndentedJSON(http.StatusConflict, err.Error())
		return
	}

	body := &responses.SignUpRequest{}
	err = json.Unmarshal(bodyBytes, body)
	if err != nil {
		common.LogError(err.Error())
		ctx.IndentedJSON(http.StatusConflict, err.Error())
		return
	}

	bodyIsValid := func() bool {
		return body.FullName != "" &&
			body.UserName != "" &&
			common.EmailIsValid(body.Email) &&
			body.Password != "" &&
			body.Age != 0 &&
			body.Role != ""
	}

	if !bodyIsValid() {
		err = errors.New("some field is empty or isn't valid")
		common.LogError(err.Error())
		ctx.IndentedJSON(http.StatusConflict, err.Error())
		return
	}

	// Redirects to GRPC

	ctx.IndentedJSON(http.StatusOK, body)
}

func HandleRefreshTwoFactorCode(ctx *gin.Context) {
	bodyBytes, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		common.LogError(err.Error())
		ctx.IndentedJSON(http.StatusConflict, err.Error())
		return
	}

	body := &responses.RefreshTwoFactorCodeRequest{}
	err = json.Unmarshal(bodyBytes, body)
	if err != nil {
		common.LogError(err.Error())
		ctx.IndentedJSON(http.StatusConflict, err.Error())
		return
	}

	if !common.EmailIsValid(body.Email) {
		err = errors.New("email isn't valid")
		common.LogError(err.Error())
		ctx.IndentedJSON(http.StatusConflict, err.Error())
		return
	}

	status, err := common.AuthServiceClient.RefreshTwoFactorCode(ctx.Request.Context(), responses.NewRefreshTwoFactorCodeRequestFromJSONToProto(body))
	if err != nil {
		common.LogError(err.Error())
		ctx.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.IndentedJSON(http.StatusOK, status.GetMessage())
}
