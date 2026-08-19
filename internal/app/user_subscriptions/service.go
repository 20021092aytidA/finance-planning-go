package usersubscriptions

import "finance-planning-go/internal/app/database"

func (us UserSubscriptionService) Get(query AllowedQuery) (error, []ViewModel) {
	var userSubs []ViewModel
	offset := (query.Page - 1) * query.Limit
	if err := database.DB.Table("user_subscriptions").Preload("User").Preload("Subscription").Where(&query.DataQuery).Offset(offset).Limit(query.Limit).Find(&userSubs).Error; err != nil {
		return err, userSubs
	}

	return nil, userSubs
}
