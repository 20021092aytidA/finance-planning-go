package users

func (ViewModel) TableName() string {
	return "users"
}

type UserHandler struct{}
type UserService struct{}
type UserRoute struct{}

type DataQuery struct {
	Id       string `form:"id"`
	Email    string `form:"email"`
	Username string `form:"username"`
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
	Id       *int    `db:"id" json:"id" gorm:"primaryKey"`
	Email    *string `db:"email" json:"email"`
	Username *string `db:"username" json:"username"`
}

type ViewModelWithPass struct {
	ViewModel
	Password *string `db:"password" json:"password"`
}
