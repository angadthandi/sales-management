package main

import (
	"fmt"
	"udemy-sales-management/item"
	"udemy-sales-management/store"
)

func main() {
	fmt.Println("sales management!")

	fmt.Println(item.Sandwich)
	fmt.Println(item.Tea)

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
	fmt.Println(store1.GetFoodSupply())
	fmt.Println(store1.GetBeverageSupply())
}
