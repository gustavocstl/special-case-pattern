package repository

import "github.com/gucastiliao/special-case-pattern/pkg/model"

type SubscriptionRepository struct{}

func (r SubscriptionRepository) GetExpiredSubscriptions() []model.Subscription {
	return []model.Subscription{
		{
			ID:               1,
			CreditCardNumber: "111",
			IsFraud:          true,
			IsIncomplete:     false,
		},
		{
			ID:               2,
			CreditCardNumber: "222",
			IsFraud:          false,
			IsIncomplete:     false,
		},
		{
			ID:               3,
			CreditCardNumber: "333",
			IsFraud:          false,
			IsIncomplete:     true,
		},
	}
}
