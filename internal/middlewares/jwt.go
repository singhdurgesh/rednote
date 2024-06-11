package middlewares

import (
	"net/http"
	"strings"

	"github.com/singhdurgesh/rednote/internal/app/services"
	"github.com/singhdurgesh/rednote/internal/pkg/utils"

	"github.com/gin-gonic/gin"
)

var userService = new(services.UserService)

// Jwt middleware
func Jwt() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		tokenStr := ctx.GetHeader("Authorization")
		if tokenStr == "" {
			ctx.Abort()
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "message": "Not Authorized"})
			return
		}
		parts := strings.Split(tokenStr, " ")

		if len(parts) != 2 || parts[0] != "Bearer" {
			ctx.Abort()
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "message": "Not Authorized"})
			return
		}

		token := parts[1]
		claims, err := utils.JwtVerify(token)

		if err != nil || claims == nil {
			ctx.Abort()
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "message": "Not Authorized"})
			return
		}

		user := userService.GetActiveUserById(int(claims.Uid))

		if user == nil {
			ctx.Abort()
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "message": "Inactive User"})
			return
		}

		ctx.Set("currentUser", user)
		ctx.Set("authProvider", claims.AuthMode)

		ctx.Next()
	}

}
