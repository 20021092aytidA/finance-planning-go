package routes

import (
	"finance-planning-go/internal/app/plans"

	"github.com/gin-gonic/gin"
)

func Init(c *gin.Engine) {
	v1 := c.Group("/finance-planning-api/v1")
	plans.InitRoute(v1)
}
