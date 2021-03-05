package payment

import (
	"github.com/gucastiliao/special-case-pattern/pkg/charge"
	"github.com/gucastiliao/special-case-pattern/pkg/model"
)

type PaymentReceiver struct{}

func (p PaymentReceiver) Charge(subscription model.Subscription) error {
	charge := charge.NewChargeFactory(subscription)
	return charge.Execute()
}
