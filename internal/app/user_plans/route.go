package userplans

import "github.com/gin-gonic/gin"

func (up UserPlanRoute) InitRoute(c *gin.RouterGroup) {
	c.GET("/user-plans", UserPlanHandler{}.View)
}
