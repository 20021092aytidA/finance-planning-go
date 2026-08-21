package plans

func (ViewModel) TableName() string {
	return "plans"
}

type PlanHandler struct{}
type PlanService struct{}
type PlanRoute struct{}

type DataQuery struct {
	Id             string `form:"id"`
	UserID         string `form:"userID"`
	Name           string `form:"name"`
	Price          string `form:"price"`
	MoneyAllocated string `form:"moneyAllocated"`
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
	Id             *int     `db:"id" json:"id" gorm:"primaryKey"`
	UserID         *int     `db:"user_id" json:"userID"`
	Name           *string  `db:"name" json:"name"`
	Price          *float32 `db:"price" json:"price"`
	MoneyAllocated *float32 `db:"money_allocated" json:"moneyAllocated"`
}

type PostModel struct {
	UserID         *int     `db:"user_id" json:"userID" form:"userID" binding:"required"`
	Name           *string  `db:"name" json:"name" form:"name" binding:"required"`
	Price          *float32 `db:"price" json:"price" form:"price" binding:"required"`
	MoneyAllocated *float32 `db:"money_allocated" json:"moneyAllocated" form:"moneyAllocated"`
}
