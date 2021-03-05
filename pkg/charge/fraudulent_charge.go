package charge

import (
	"errors"

	"github.com/gucastiliao/special-case-pattern/pkg/model"
)

type FraudulentCharge struct {
	subscription model.Subscription
}

func NewFraudulentCharge(subscription model.Subscription) FraudulentCharge {
	return FraudulentCharge{
		subscription: subscription,
	}
}

func (c FraudulentCharge) Execute() error {
	c.setFraudulentCharge()
	return errors.New("Fraudulent charge")
}

func (c FraudulentCharge) setFraudulentCharge() error {
	return nil
}
