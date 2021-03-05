package charge

import (
	"errors"

	"github.com/gucastiliao/special-case-pattern/pkg/model"
)

type IncompleteCharge struct {
	subscription model.Subscription
}

func NewIncompleteCharge(subscription model.Subscription) IncompleteCharge {
	return IncompleteCharge{
		subscription: subscription,
	}
}

func (c IncompleteCharge) Execute() error {
	c.setIncompleteCharge()
	return errors.New("Incomplete charge")
}

func (c IncompleteCharge) setIncompleteCharge() error {
	return nil
}
