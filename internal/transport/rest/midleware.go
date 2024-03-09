package rest

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Ilyasich/weather/internal/services"
	"github.com/gin-gonic/gin"
)

func AuthMiddlevare(service *services.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenAuth := ctx.GetHeader("Authorization")
		if tokenAuth == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token is missing"})
			ctx.Abort()
			return
		}
		if !strings.HasPrefix(tokenAuth, "Token") {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error":"Invalid token format"})
			return
		}
		token := strings.TrimPrefix(tokenAuth, "Token")

		userJSON, err := base64.StdEncoding.DecodeString(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			ctx.Abort()
			return
		}

		var username string
		if err := json.Unmarshal(userJSON, &username); err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			ctx.Abort()
			return
		}

		exist, err := service.UserExists(ctx, username)
		if err != nil || !exist {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error":"Invalid token"})
			return
		}
		ctx.Set("user", username)
		ctx.Next()

	}

}
