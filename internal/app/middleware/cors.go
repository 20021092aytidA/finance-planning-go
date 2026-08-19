package middleware

import (
	"finance-planning-go/internal/app/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type CORS struct{}

func (c CORS) Use() gin.HandlerFunc {
	return cors.New(config.CORS{}.Setting())
}
