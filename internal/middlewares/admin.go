package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/singhdurgesh/rednote/internal/app/models"
	"github.com/singhdurgesh/rednote/internal/constants"
)

func AdminAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userData, _ := ctx.Get("currentUser")
		user := userData.(*models.User)
		authProviderData, _ := ctx.Get("authProvider")
		if user.EmailVerified && authProviderData == constants.GOOGLE_AUTH_MODE {
			ctx.Next()
		} else {
			ctx.Abort()
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "message": "User is Not Admin"})
			return
		}
	}
}
