package main

import (
	"fmt"

	"github.com/gucastiliao/special-case-pattern/pkg/payment"
	"github.com/gucastiliao/special-case-pattern/pkg/repository"
)

func main() {
	subscriptionRepository := repository.SubscriptionRepository{}
	paymentReceiver := payment.PaymentReceiver{}

	expiredSubscriptions := subscriptionRepository.GetExpiredSubscriptions()

	for _, subscription := range expiredSubscriptions {
		fmt.Println(paymentReceiver.Charge(subscription))
	}
}
