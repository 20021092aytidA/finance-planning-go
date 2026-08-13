package cors_cfg

import (
	"time"

	"github.com/gin-contrib/cors"
)

func Load() cors.Config {
	corsCfg := cors.Config{
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "X-CSRF-Token"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: true,
		MaxAge:           time.Hour * 1,
	}

	return corsCfg
}
