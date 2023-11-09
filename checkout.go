package main

import "fmt"

type CheckoutFacade struct {
    shop *Shop
}

func NewCheckoutFacade(shop *Shop) *CheckoutFacade {
    return &CheckoutFacade{shop: shop}
}

func (cf *CheckoutFacade) CheckoutWithPayment(productID int, paymentStrategy PaymentStrategy) {
    product := cf.shop.Inventory.GetProductByID(productID)
    if product == nil {
        fmt.Println("Product not found.")
        return
    }

    if !paymentStrategy.Pay(product.Price) {
        fmt.Println("Payment failed.")
        return
    }

    cf.shop.funds -= product.Price
    cf.shop.purchasedProducts = append(cf.shop.purchasedProducts, *product)
    fmt.Println("Product purchased:", product.Name)
}
