package router

import (
	"github.com/singhdurgesh/rednote/internal/app/controllers/webhookControllers"

	"github.com/gin-gonic/gin"
)

var webhookController = new(webhookControllers.OAuthController)

func LoadWebhookRoutes(r *gin.Engine) *gin.RouterGroup {

	webhooks := r.Group("/webhooks")

	oauth := webhooks.Group("/oauth")
	{
		oauth.GET("/googleLogin", webhookController.GoogleLogin)
	}

	return oauth
}
