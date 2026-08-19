package plans

type DataQuery struct {
	Id             string `form:"id"`
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
	Id             *int     `db:"id" json:"id"`
	Name           *string  `db:"name" json:"name"`
	Price          *float32 `db:"price" json:"price"`
	MoneyAllocated *float32 `db:"money_allocated" json:"moneyAllocated"`
}
