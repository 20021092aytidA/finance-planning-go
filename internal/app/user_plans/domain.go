package userplans

import (
	"finance-planning-go/internal/app/plans"
	"finance-planning-go/internal/app/users"
)

func (ViewModel) TableName() string {
	return "user_plans"
}

type UserPlanHandler struct{}
type UserPlanService struct{}
type UserPlanRoute struct{}

type DataQuery struct {
	Id     string `form:"id"`
	UserID string `form:"userID"`
	PlanID string `form:"planID"`
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
	Id     *int `db:"id" json:"id" gorm:"primaryKey"`
	UserID *int `db:"user_id" json:"userID"`
	PlanID *int `db:"plan_id" json:"planID"`

	User users.ViewModel `json:"user" gorm:"foreignKey:UserID;references:Id"`
	Plan plans.ViewModel `json:"plan" gorm:"foreignKey:PlanID;references:Id"`
}
