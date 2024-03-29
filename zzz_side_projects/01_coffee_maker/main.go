package main

import "fmt"

func main() {
	var isOn = true

	var water int = 4000
	var milk int = 4000
	var coffee int = 2000
	var money float32 = 0

	for isOn {
		var choice string
		fmt.Println("What coffee you would like to order?\nEspresso, Cappuccino or Latte:")
		fmt.Scanln(&choice)
		fmt.Println("You choose", choice)

		switch choice {
		case "off":
			{
				fmt.Println("Good bye...")
				isOn = false
			}
		case "report":
			{
				fmt.Println("Water:", water, "ml")
				fmt.Println("Milk:", milk, "ml")
				fmt.Println("Coffee:", coffee, "ml")
				fmt.Println("Money:", money, "€")
			}
		case "espresso":
			{
				if water < 50 || coffee < 18 {
					fmt.Println("Sorry, this drink is not available at the moment. Please make a new order.")
					break
				} else {
					// TODO 2 - price 1.5
					drinkPaid, cashOut := cashier(1.5)
					fmt.Println("Thanks for your order, your change is", cashOut, "€.")
					fmt.Println("Enjoy your", choice, ".")

					if drinkPaid {
						money += 1.5
						water -= 50
						coffee -= 18
					}
				}
			}
		case "cappuccino":
			{
				if water < 250 || milk < 100 || coffee < 24 {
					fmt.Println("Sorry, this drink is not available at the moment. Please make a new order.")
					break
				} else {
					// TODO 3 - price 3
					drinkPaid, cashOut := cashier(3)
					fmt.Println("Thanks for your order, your change is", cashOut, "€.")
					fmt.Println("Enjoy your", choice, ".")

					if drinkPaid {
						money += 3
						water -= 250
						milk -= 100
						coffee -= 24
					}
				}
			}
		case "latte":
			{
				if water < 200 || milk < 150 || coffee < 24 {
					fmt.Println("Sorry, this drink is not available at the moment. Please make a new order.")
					break
				} else {
					// TODO 4 - price 2.5
					drinkPaid, cashOut := cashier(2.5)
					fmt.Println("Thanks for your order, your change is", cashOut, "€.")
					fmt.Println("Enjoy your", choice, ".")

					if drinkPaid {
						money += 2.5
						water -= 200
						milk -= 150
						coffee -= 24
					}
				}
			}
		default:
			fmt.Println("Sorry, this is not available. Please make a new order.")
		}
	}

}

// TODO 1
func cashier(price float32) (bool, float32) {
	var quarters float32
	var dimes float32
	var nickles float32
	var pennies float32

	fmt.Println("Please insert quarters")
	fmt.Scanln(&quarters)
	fmt.Println("Please insert dimes")
	fmt.Scanln(&dimes)
	fmt.Println("Please insert nickles")
	fmt.Scanln(&nickles)
	fmt.Println("Please insert pennies")
	fmt.Scanln(&pennies)

	var cashIn float32 = quarters*0.25 + dimes*0.1 + nickles*0.05 + pennies*0.01
	var drinkPaid bool = false

	if cashIn >= price {
		drinkPaid = true
	}

	var cashOut float32 = cashIn - price

	return drinkPaid, cashOut
}
