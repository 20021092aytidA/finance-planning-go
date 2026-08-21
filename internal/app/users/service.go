package users

import "finance-planning-go/internal/app/database"

func (u UserService) Get(query *AllowedQuery) (error, []ViewModel) {
	var users []ViewModel
	offset := (query.Page - 1) * query.Limit
	if err := database.DB.Table("users").Where(query.DataQuery).Offset(offset).Limit(query.Limit).Find(&users).Error; err != nil {
		return err, users
	}

	return nil, users
}

func (u UserService) Post(newUser *PostModel) error {
	if err := database.DB.Table("users").Create(newUser).Error; err != nil {
		return err
	}
	return nil
}
