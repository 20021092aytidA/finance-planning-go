package routes

import (
	"finance-planning-go/internal/app/plans"
	"finance-planning-go/internal/app/subscriptions"
	userplans "finance-planning-go/internal/app/user_plans"
	usersubscriptions "finance-planning-go/internal/app/user_subscriptions"
	"finance-planning-go/internal/app/users"

	"github.com/gin-gonic/gin"
)

func Init(c *gin.Engine) {
	v1 := c.Group("/finance-planning-api/v1")

	plans.PlanRoute{}.InitRoute(v1)
	subscriptions.SubscriptionRoute{}.InitRoute(v1)
	users.UserRoute{}.InitRoute(v1)
	userplans.UserPlanRoute{}.InitRoute(v1)
	usersubscriptions.UserSubscriptionRoute{}.InitRoute(v1)
}
