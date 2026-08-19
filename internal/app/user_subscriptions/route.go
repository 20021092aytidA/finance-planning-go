package usersubscriptions

import "github.com/gin-gonic/gin"

func (us UserSubscriptionRoute) InitRoute(c *gin.RouterGroup) {
	c.GET("/user_subscriptions", UserSubscriptionHanlder{}.View)
}
