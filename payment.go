package main

import "fmt"

// PaymentStrategy interface
type PaymentStrategy interface {
    Pay(amount float64) bool
}

// CashPayment strategy
type CashPayment struct{}

func (cp *CashPayment) Pay(amount float64) bool {
    fmt.Printf("Paid $%.2f using cash.\n", amount)
    return true // Assuming cash payment always succeeds
}

// PayPalPayment strategy
type PayPalPayment struct {
    email string
}

func (pp *PayPalPayment) Pay(amount float64) bool {
    fmt.Printf("Paid $%.2f using PayPal (Email: %s).\n", amount, pp.email)
    return true // Assuming PayPal payment always succeeds
}
