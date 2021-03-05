package charge

import "github.com/gucastiliao/special-case-pattern/pkg/model"

type CompleteCharge struct {
	subscription model.Subscription
}

func NewCompleteCharge(subscription model.Subscription) CompleteCharge {
	return CompleteCharge{
		subscription: subscription,
	}
}

func (c CompleteCharge) Execute() error {
	c.setCompleteCharge()
	return nil
}

func (c CompleteCharge) setCompleteCharge() error {
	return nil
}
