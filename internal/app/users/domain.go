package users

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (ViewModel) TableName() string {
	return "users"
}

// HASH PASSWORD
func (u *PostModel) BeforeCreate(tx *gorm.DB) error {
	if errHash := hashPassword(u); errHash != nil {
		return errHash
	}

	return nil
}

func hashPassword(u *PostModel) error {
	byteHash, errHash := bcrypt.GenerateFromPassword([]byte(*u.Password), bcrypt.DefaultCost)
	if errHash != nil {
		return errHash
	}

	stringHash := string(byteHash[:])
	u.Password = &stringHash
	return nil
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

type PostModel struct {
	Id       *int    `db:"id" json:"id" gorm:"primaryKey"`
	Email    *string `db:"email" json:"email" form:"email" binding:"required"`
	Username *string `db:"username" json:"username" form:"username" binding:"required"`
	Password *string `db:"password" json:"password" form:"password" binding:"required"`
}
