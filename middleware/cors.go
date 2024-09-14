package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	corsconfig := cors.DefaultConfig()
	corsconfig.AllowCredentials = true
	corsconfig.AllowOrigins = []string{"*"}
	corsconfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	return cors.New(corsconfig)
}
