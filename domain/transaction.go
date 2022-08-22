package domain

import "time"

type TransactionRepository interface {
	SaveTransaction(transaction Transaction, creditCard CreditCard) error
	GetCreditCard(creditCard CreditCard) (CreditCard, error)
	CreateCreditCard(creditCard CreditCard) error
}

type Transaction struct {
	ID string
	Amount string
	Status string
	Description string
	Store string
	CreditCardId string
	CreatedAt	time.Time
}

func NewTransaction() *Transaction{
	t := &Transaction{}
	t.ID = uuid.NewV4().String()
	c.CreatedAt = time.now()
	return c
}

func (t *Transaction) ProcessAndValidate(creditCard *CreditCard) {
	if t.Amount + creditCard.Balance > creditCard.Limit {
		t.Status = "rejected"
	}	else {
		t.Status = "approved"
		creditCard.Balance = creditCard.Balance + t.Amount
	}
}