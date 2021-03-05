package model

type Subscription struct {
	ID               int
	CreditCardNumber string
	IsFraud          bool
	IsIncomplete     bool
}
