package users

import "github.com/gin-gonic/gin"

func (u UserRoute) InitRoute(c *gin.RouterGroup) {
	c.GET("/users", UserHandler{}.View)
	c.POST("/users", UserHandler{}.Create)
}
