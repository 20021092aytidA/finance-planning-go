package subscriptions

import "finance-planning-go/internal/app/database"

func (s SubscriptionService) Get(query AllowedQuery) (error, []ViewModel) {
	var subs []ViewModel
	offset := (query.Page - 1) * query.Limit
	if err := database.DB.Table("subscriptions").Where(&query.DataQuery).Offset(offset).Limit(query.Limit).Find(&subs).Error; err != nil {
		return err, subs
	}

	return nil, subs
}
