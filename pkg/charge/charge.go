package charge

import "github.com/gucastiliao/special-case-pattern/pkg/model"

type Charge interface {
	Execute() error
}

func NewChargeFactory(subscription model.Subscription) Charge {
	if isFraudDetected(subscription) {
		return NewFraudulentCharge(subscription)
	}

	if isSubscriptionIncomplete(subscription) {
		return NewIncompleteCharge(subscription)
	}

	return NewCompleteCharge(subscription)
}

func isFraudDetected(subscription model.Subscription) bool {
	return subscription.IsFraud
}

func isSubscriptionIncomplete(subscription model.Subscription) bool {
	return subscription.IsIncomplete
}
