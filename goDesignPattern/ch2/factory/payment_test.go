package factory

import (
	"strings"
	"testing"
)

func TestCreatePaymentMethodCash(t *testing.T) {
	payment, err := GetPaymentMethod(Cash)
	if err != nil {
		t.Errorf("A payment of type 'Cash' must exist")
	}

	msg := payment.Pay(10.43)
	if !strings.Contains(msg, "paid using cash") {
		t.Errorf("The cash payment method message was not correct")
	}
	t.Log("LOG:", msg)
}

func TestCreatePaymentMethodDebitCard(t *testing.T) {
	payment, err := GetPaymentMethod(DebitCard)
	if err != nil {
		t.Errorf("A payment of type 'DebitCard' must exist")
	}

	msg := payment.Pay(12.43)
	if !strings.Contains(msg, "paid using debit card") {
		t.Errorf("The cash payment method message was not correct")
	}
	t.Log("LOG:", msg)
}

func TestCreatePaymentMethodNonExistent(t *testing.T) {
	_, err := GetPaymentMethod(20)
	if err == nil {
		t.Errorf("A payment with ID 20 must return error")
	}

	t.Log("LOG:", err)
}
