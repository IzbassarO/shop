package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	shop := GetShopInstance()

	fmt.Println("Welcome to the Go Shop!")
	fmt.Print("Please enter your name: ")
	name, _ := reader.ReadString('\n')
	fmt.Print("Hello, ", strings.TrimSpace(name), "! How much money do you have in cash? (Minimum $10): ")
	fundsStr, _ := reader.ReadString('\n')
	fundsStr = strings.TrimSpace(fundsStr)
	funds, err := strconv.ParseFloat(fundsStr, 64)
	if err != nil || funds < 10 {
		fmt.Println("Invalid amount of money. Setting default to $100.")
		funds = 100
	}
	shop.AddFunds(funds)
	fmt.Println("For first register, you will get $100. Congratulations!")

	for {
		fmt.Println("Select an option:")
		fmt.Println("1. List products")
		fmt.Println("2. Check balance")
		fmt.Println("3. Check purchased products")
		fmt.Println("4. Quit")
		fmt.Print("Enter option: ")
		optionStr, _ := reader.ReadString('\n')
		optionStr = strings.TrimSpace(optionStr)
		option, err := strconv.Atoi(optionStr)
		if err != nil {
			fmt.Println("Invalid option.")
			continue
		}

		switch option {
		case 1:
			shop.Inventory.ListProducts()
			fmt.Print("Enter product ID to purchase or 'back' to return to main menu: ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			if input == "back" {
				continue
			}

			productID, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("Invalid product ID.")
				continue
			}

			shop.CheckoutFacade.Checkout(productID)

		case 2:
			fmt.Printf("Your current balance is $%.2f\n", shop.GetFunds())

		case 3:
			purchasedProducts := shop.GetPurchasedProducts()
			if len(purchasedProducts) == 0 {
				fmt.Println("You haven't purchased any products yet.")
			} else {
				fmt.Println("Purchased Products:")
				for _, p := range purchasedProducts {
					fmt.Printf("- %s: $%.2f\n", p.Name, p.Price)
				}
			}

		case 4:
			fmt.Println("Thank you for shopping with us, goodbye!")
			return

		default:
			fmt.Println("Option not recognized, please try again.")
		}
	}
}
