package apiserver

import (
	"net/http"

	model "apiserver/v1/nomad-model"
	utils "apiserver/v1/nomad-utils"

	"github.com/gin-gonic/gin"
)

func AuthorizationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		api_resp := model.ApiResponse{
			Code:    401,
			Data:    model.ApiResponseData{},
			Message: "No authorization token provided",
		}

		var authorization string
		if c.Query("Authorization") != "" {
			authorization = c.Query("Authorization")
		} else {
			authorization = c.Query("authorization")
		}
		if authorization == "" {
			api_resp.Message = "Authorization key is required"
			c.JSON(
				http.StatusUnauthorized,
				api_resp,
			)
			c.Abort()
			return
		}

		if authorization != utils.AppConfig().SecurityCaKey {
			api_resp.Message = "Provided apikey is not valid"
			c.JSON(
				http.StatusUnauthorized,
				api_resp,
			)
			c.Abort()
			return
		}

		c.Next()
	}
}
