package main

import (
	"fmt"
)

type CheckoutFacade struct {
	shop *Shop
}

func NewCheckoutFacade(shop *Shop) *CheckoutFacade {
	return &CheckoutFacade{shop: shop}
}

func (cf *CheckoutFacade) Checkout(productID int) {
	product := cf.shop.Inventory.GetProductByID(productID)
	if product == nil {
		fmt.Println("Product not found.")
		return
	}

	if cf.shop.funds < product.Price {
		fmt.Println("Insufficient funds.")
		return
	}

	cf.shop.funds -= product.Price
	cf.shop.purchasedProducts = append(cf.shop.purchasedProducts, *product)
	fmt.Println("Product purchased:", product.Name)
}
