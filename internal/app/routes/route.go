package routes

import (
	"finance-planning-go/internal/app/plans"
	"finance-planning-go/internal/app/subscriptions"
	"finance-planning-go/internal/app/users"

	"github.com/gin-gonic/gin"
)

func Init(c *gin.Engine) {
	v1 := c.Group("/finance-planning-api/v1")

	plans.PlanRoute{}.InitRoute(v1)
	subscriptions.SubscriptionRoute{}.InitRoute(v1)
	users.UserRoute{}.InitRoute(v1)
}
