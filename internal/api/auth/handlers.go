package auth

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/carlosabdoamaral/behealthy_api/common"
	"github.com/gin-gonic/gin"
)

func HandleSignIn(ctx *gin.Context) {
	bodyBytes, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		common.LogError(err.Error())
		ctx.IndentedJSON(http.StatusConflict, err.Error())
		return
	}

	body := &SignInRequestModel{}
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

	// Redirects to GRPC

	ctx.IndentedJSON(http.StatusOK, body)
}

func HandleSignUp(ctx *gin.Context) {
	bodyBytes, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		common.LogError(err.Error())
		ctx.IndentedJSON(http.StatusConflict, err.Error())
		return
	}

	body := &SignUpRequestModel{}
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

func HandleRecoverPassword(ctx *gin.Context) {
	bodyBytes, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		common.LogError(err.Error())
		ctx.IndentedJSON(http.StatusConflict, err.Error())
		return
	}

	body := &RecoverPasswordRequestModel{}
	err = json.Unmarshal(bodyBytes, body)
	if err != nil {
		common.LogError(err.Error())
		ctx.IndentedJSON(http.StatusConflict, err.Error())
		return
	}

	bodyIsValid := func() bool {
		return common.EmailIsValid(body.Email)
	}

	if !bodyIsValid() {
		err = errors.New("email isn't valid")
		common.LogError(err.Error())
		ctx.IndentedJSON(http.StatusConflict, err.Error())
		return
	}

	// Redirects to GRPC

	ctx.IndentedJSON(http.StatusOK, body)
}

func HandleRecoverPasswordValidation(ctx *gin.Context) {
	bodyBytes, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		common.LogError(err.Error())
		ctx.IndentedJSON(http.StatusConflict, err.Error())
		return
	}

	body := &RecoverPasswordValidationRequestModel{}
	err = json.Unmarshal(bodyBytes, body)
	if err != nil {
		common.LogError(err.Error())
		ctx.IndentedJSON(http.StatusConflict, err.Error())
		return
	}

	bodyIsValid := func() bool {
		return body.TwoFactorCode != "" && body.NewPassword != ""
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
