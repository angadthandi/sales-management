package main

import (
	"fmt"
	"udemy-sales-management/city"
	"udemy-sales-management/item"
	"udemy-sales-management/state"
	"udemy-sales-management/store"
	"udemy-sales-management/system"
)

func main() {
	fmt.Println("sales management!")

	foodSupplyMap := make(map[item.FoodItem]int)
	foodSupplyMap[item.Sandwich] = 5
	foodSupplyMap[item.Burger] = 5
	foodSupplyMap[item.Pizza] = 5
	foodSupplyMap[item.Noodles] = 5

	beverageSupplyMap := make(map[item.BeverageItem]int)
	beverageSupplyMap[item.Tea] = 5
	beverageSupplyMap[item.Coffee] = 5
	beverageSupplyMap[item.Water] = 5

	store1 := store.GetInstance(foodSupplyMap, beverageSupplyMap)
	// fmt.Println(store1.GetFoodSupply())
	// fmt.Println(store1.GetBeverageSupply())

	city1 := city.GetInstance(foodSupplyMap, beverageSupplyMap)
	city1.AddStore(store1)

	state1 := state.GetInstance()
	state1.AddCity(city1)

	system1 := system.GetInstance()
	system1.AddState(state1)

	system1.PurchaseFood(
		state1.GetID(),
		city1.GetID(),
		store1.GetID(),
		item.Burger,
		2,
	)

	for _, state := range system1.GetStates() {

		if state.GetID() == state1.GetID() {
			for _, city := range state.GetCities() {

				if city.GetID() == city1.GetID() {
					for _, store := range city.GetStores() {

						if store.GetID() == store1.GetID() {

							for food, sold := range store.GetFoodUnitsSold() {

								fmt.Printf("%v : %v\n", food, sold)

							}
							break

						}

					}
					break
				}

			}
			break
		}

	}

	// should fail
	// total=5 burgers
	// already purchased=2 burgers
	// remaining=3 burgers
	system1.PurchaseFood(
		state1.GetID(),
		city1.GetID(),
		store1.GetID(),
		item.Burger,
		4,
	)
}
