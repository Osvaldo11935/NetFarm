package middleware

import "github.com/gin-gonic/gin"

func SetHeaderMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/json")
	}
}
