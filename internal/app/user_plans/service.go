package userplans

import "finance-planning-go/internal/app/database"

func (up UserPlanService) Get(query AllowedQuery) (error, []ViewModel) {
	var userPlans []ViewModel
	offset := (query.Page - 1) * query.Limit
	if err := database.DB.Table("user_plans").Preload("User").Preload("Plan").Where(&query.DataQuery).Offset(offset).Limit(query.Limit).Find(&userPlans).Error; err != nil {
		return err, userPlans
	}

	return nil, userPlans
}
