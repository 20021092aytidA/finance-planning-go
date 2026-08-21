package subscriptions

import "github.com/gin-gonic/gin"

func (s SubscriptionRoute) InitRoute(c *gin.RouterGroup) {
	c.GET("/subscriptions", SubscriptionHandler{}.View)
	c.POST("/subscriptions", SubscriptionHandler{}.Create)

	c.DELETE("/subscriptions/:id", SubscriptionHandler{}.Drop)
}
