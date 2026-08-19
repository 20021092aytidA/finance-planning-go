package subscriptions

type SubscriptionService struct{}
type SubscriptionHandler struct{}
type SubscriptionRoute struct{}

type DataQuery struct {
	Id        string `form:"id"`
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
	Id        *int     `db:"id" json:"id"`
	Name      *string  `db:"name" json:"name"`
	Interval  *int     `db:"interval" json:"interval"`
	Price     *float32 `db:"price" json:"price"`
	StartDate *string  `db:"start_date" json:"startDate"`
}
