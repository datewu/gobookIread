package factory

import "fmt"

// PaymentMethod interface define the way of making payment
type PaymentMethod interface {
	Pay(amount float32) string
}

// Cash payment
const (
	Cash      = 1
	DebitCard = 2
)

// GetPaymentMethod return a pointer implements the interface
func GetPaymentMethod(m int) (PaymentMethod, error) {
	switch m {
	case Cash:
		return new(CashPM), nil
	case DebitCard:
		return new(DebitCardPM), nil
	default:
		return nil, fmt.Errorf("payment method %d not recognized", m)
	}
}

// CashPM struct
type CashPM struct{}

// DebitCardPM struct
type DebitCardPM struct{}

// Pay satisfy PaymentMethod interface
func (c *CashPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using cash\n", amount)
}

// Pay satisfy PaymentMethod interface
func (d *DebitCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using debit card\n", amount)
}
