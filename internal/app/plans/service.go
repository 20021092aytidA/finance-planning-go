package plans

import "finance-planning-go/internal/app/database"

func (p PlanService) Get(query AllowedQuery) (error, []ViewModel) {
	var plans []ViewModel
	offset := (query.Page - 1) * query.Limit
	if err := database.DB.Table("plans").Where(&query.DataQuery).Offset(offset).Limit(query.Limit).Find(&plans).Error; err != nil {
		return err, plans
	}

	return nil, plans
}
