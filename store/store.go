package store

import (
	"fmt"
	"strconv"
	"udemy-sales-management/item"
)

type IStore interface {
	GetID() string
	GetFoodSupply() map[item.FoodItem]int
	GetBeverageSupply() map[item.BeverageItem]int
	GetFoodUnitsSold() map[item.FoodItem]int
	GetBeverageUnitsSold() map[item.BeverageItem]int
	SetFoodRates(rates map[item.FoodItem]int)
	SetBeverageRates(rates map[item.BeverageItem]int)
	PurchaseFood(f item.FoodItem, qty int)
	PurchaseBeverage(b item.BeverageItem, qty int)
}

type Store struct {
	ID                string
	FoodSupply        map[item.FoodItem]int
	BeverageSupply    map[item.BeverageItem]int
	FoodUnitsSold     map[item.FoodItem]int
	BeverageUnitsSold map[item.BeverageItem]int
	FoodRates         map[item.FoodItem]int
	BeverageRates     map[item.BeverageItem]int
}

var idIncrementor int

func GetInstance(
	foodSupplyMap map[item.FoodItem]int, // {item: qty}
	beverageSupplyMap map[item.BeverageItem]int, // {item: qty}
) IStore {
	idIncrementor++

	s := Store{
		ID:                "store_" + strconv.Itoa(idIncrementor),
		FoodSupply:        foodSupplyMap,
		BeverageSupply:    beverageSupplyMap,
		FoodUnitsSold:     make(map[item.FoodItem]int),
		BeverageUnitsSold: make(map[item.BeverageItem]int),
		FoodRates:         make(map[item.FoodItem]int),
		BeverageRates:     make(map[item.BeverageItem]int),
	}

	return s
}

func (s Store) GetID() string {
	return s.ID
}

func (s Store) GetFoodSupply() map[item.FoodItem]int {
	return s.FoodSupply
}

func (s Store) GetBeverageSupply() map[item.BeverageItem]int {
	return s.BeverageSupply
}

func (s Store) GetFoodUnitsSold() map[item.FoodItem]int {
	return s.FoodUnitsSold
}

func (s Store) GetBeverageUnitsSold() map[item.BeverageItem]int {
	return s.BeverageUnitsSold
}

func (s Store) SetFoodRates(rates map[item.FoodItem]int) {
	s.FoodRates = rates
}

func (s Store) SetBeverageRates(rates map[item.BeverageItem]int) {
	s.BeverageRates = rates
}

func (s Store) PurchaseFood(f item.FoodItem, qty int) {
	if _, ok := s.FoodSupply[f]; !ok {
		fmt.Printf("not enough food item: %v", f.GetItem())
	}

	s.FoodSupply[f] -= 1
	s.FoodUnitsSold[f] += 1
}

func (s Store) PurchaseBeverage(b item.BeverageItem, qty int) {
	if _, ok := s.BeverageSupply[b]; !ok {
		fmt.Printf("not enough beverage item: %v", b.GetItem())
	}

	s.BeverageSupply[b] -= 1
	s.BeverageUnitsSold[b] += 1
}
