package main

import "fmt"

// PaymentStrategy interface defines how the payment should be processed.
type PaymentStrategy interface {
	Pay(amount float64) bool
}

// CashPayment is a Strategy for paying with cash.
type CashPayment struct{}

func (cp *CashPayment) Pay(amount float64) bool {
	fmt.Printf("Paid $%.2f using cash.\n", amount)
	return true // Assuming cash payment always succeeds
}