package city

import (
	"strconv"
	"udemy-sales-management/item"
	"udemy-sales-management/store"
)

type ICity interface {
	GetID() string
	GetStores() []store.IStore
	AddStore(s store.IStore)
	PurchaseFood(storeID string, f item.FoodItem, qty int)
	PurchaseBeverage(storeID string, b item.BeverageItem, qty int)
}

type City struct {
	ID             string
	FoodPrices     map[item.FoodItem]int
	BeveragePrices map[item.BeverageItem]int
	Stores         []store.IStore
}

var idIncrementor int

func GetInstance(
	foodPriceMap map[item.FoodItem]int, // {item: price}
	beveragePriceMap map[item.BeverageItem]int, // {item: price}
) ICity {
	idIncrementor++

	c := City{
		ID:             "city_" + strconv.Itoa(idIncrementor),
		FoodPrices:     foodPriceMap,
		BeveragePrices: beveragePriceMap,
	}

	return c
}

func (c City) GetID() string {
	return c.ID
}

func (c City) GetStores() []store.IStore {
	return c.Stores
}

func (c City) AddStore(s store.IStore) {
	s.SetFoodRates(c.FoodPrices)
	s.SetBeverageRates(c.BeveragePrices)
	c.Stores = append(c.Stores, s)
}

func (c City) PurchaseFood(storeID string, f item.FoodItem, qty int) {
	for i := 0; i < len(c.Stores); i++ {
		if c.Stores[i].GetID() == storeID {
			c.Stores[i].PurchaseFood(f, qty)
			break
		}
	}
}

func (c City) PurchaseBeverage(storeID string, b item.BeverageItem, qty int) {
	for i := 0; i < len(c.Stores); i++ {
		if c.Stores[i].GetID() == storeID {
			c.Stores[i].PurchaseBeverage(b, qty)
			break
		}
	}
}
