package main

type Shop struct {
	funds             float64
	Inventory         *InventoryManager
	purchasedProducts []Product
	CheckoutFacade    *CheckoutFacade
}

func GetShopInstance() *Shop {
	shop := &Shop{
		funds:     100, // Starting funds
		Inventory: NewInventoryManager(),
	}
	shop.CheckoutFacade = NewCheckoutFacade(shop)
	return shop
}

func (s *Shop) AddFunds(funds float64) {
	s.funds += funds
}

func (s *Shop) GetFunds() float64 {
	return s.funds
}

func (s *Shop) GetPurchasedProducts() []Product {
	return s.purchasedProducts
}
