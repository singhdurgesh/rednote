package adminControllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/singhdurgesh/rednote/cmd/app"
	"github.com/singhdurgesh/rednote/pkg/oauth/googleOAuth"
)

type AuthController struct{}

func (authController *AuthController) GoogleLogin(ctx *gin.Context) {
	getUrl, err := googleOAuth.GetSSOUrl()

	if err != nil {
		app.Logger.Error("Google SSO Failure: ", err.Error())
		ctx.JSON(http.StatusInternalServerError, "Unexpected Error")
	}

	ctx.JSON(http.StatusTemporaryRedirect, getUrl)
}
