package plans

import "github.com/gin-gonic/gin"

func InitRoute(c *gin.RouterGroup) {
	c.GET("/plans", View)
}
