package subscriptions

func (ViewModel) TableName() string {
	return "subscriptions"
}

type SubscriptionService struct{}
type SubscriptionHandler struct{}
type SubscriptionRoute struct{}

type DataQuery struct {
	Id        string `form:"id"`
	UserID    string `form:"userID"`
	Name      string `form:"name"`
	Interval  string `form:"interval"`
	Price     string `form:"price"`
	StartDate string `form:"startDate"`
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
	Id        *int     `db:"id" json:"id" gorm:"primaryKey"`
	UserID    *int     `db:"user_id" json:"userID"`
	Name      *string  `db:"name" json:"name"`
	Interval  *int     `db:"interval" json:"interval"`
	Price     *float32 `db:"price" json:"price"`
	StartDate *string  `db:"start_date" json:"startDate"`
}

type PostModel struct {
	UserID    *int     `db:"user_id" json:"userID" form:"userID" binding:"required"`
	Name      *string  `db:"name" json:"name" form:"name" binding:"required"`
	Interval  *int     `db:"interval" json:"interval" form:"interval" binding:"required"`
	Price     *float32 `db:"price" json:"price" form:"price" binding:"required"`
	StartDate *string  `db:"start_date" json:"startDate" form:"startDate" binding:"required"`
}
