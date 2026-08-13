package cors_middleware

import (
	"finance-planning-go/config/cors_cfg"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func UseCORS() gin.HandlerFunc {
	return cors.New(cors_cfg.Load())
}
