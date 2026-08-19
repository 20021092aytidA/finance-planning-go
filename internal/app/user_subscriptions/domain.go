package usersubscriptions

import (
	"finance-planning-go/internal/app/subscriptions"
	"finance-planning-go/internal/app/users"
)

func (ViewModel) TableName() string {
	return "user_subscriptions"
}

type UserSubscriptionHanlder struct{}
type UserSubscriptionService struct{}
type UserSubscriptionRoute struct{}

type DataQuery struct {
	Id             string `form:"id"`
	UserID         string `form:"userID"`
	SubscriptionID string `form:"subscriptionID"`
}

type PageQuery struct {
	Page  int `form:"page"`
	Limit int `form:"limit"`
}

type AllowedQuery struct {
	DataQuery
	PageQuery
}

type ViewModel struct {
	Id             *int `db:"id" json:"id" gorm:"primaryKey"`
	UserID         *int `db:"user_id" json:"userID"`
	SubscriptionID *int `db:"subscription_id" json:"subscriptionID"`

	User         users.ViewModel         `json:"user" gorm:"foreignKey:UserID;references:Id"`
	Subscription subscriptions.ViewModel `json:"subscrpition" gorm:"foreignKey:SubscriptionID;references:Id"`
}
