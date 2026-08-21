package plans

import "github.com/gin-gonic/gin"

func (p PlanRoute) InitRoute(c *gin.RouterGroup) {
	c.GET("/plans", PlanHandler{}.View)
	c.POST("/plans", PlanHandler{}.Create)
}
