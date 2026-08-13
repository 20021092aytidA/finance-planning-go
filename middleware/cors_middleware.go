package middleware

import (
	"finance-planning-go/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func UseCORS() gin.HandlerFunc {
	return cors.New(config.LoadCORSCfg())
}
