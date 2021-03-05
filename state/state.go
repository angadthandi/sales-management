package state

import (
	"strconv"
	"udemy-sales-management/city"
	"udemy-sales-management/item"
)

type IState interface {
	GetID() string
	GetCities() []city.ICity
	AddCity(c city.ICity)
	PurchaseFood(cityID string, storeID string, f item.FoodItem, qty int)
	PurchaseBeverage(cityID string, storeID string, b item.BeverageItem, qty int)
}

type State struct {
	ID     string
	Cities []city.ICity
}

var idIncrementor int

func GetInstance() IState {
	idIncrementor++

	s := State{ID: "state_" + strconv.Itoa(idIncrementor)}

	return &s
}

func (s *State) GetID() string {
	return s.ID
}

func (s *State) GetCities() []city.ICity {
	return s.Cities
}

func (s *State) AddCity(c city.ICity) {
	s.Cities = append(s.Cities, c)
}

func (s *State) PurchaseFood(
	cityID string, storeID string, f item.FoodItem, qty int,
) {
	for i := 0; i < len(s.Cities); i++ {
		if s.Cities[i].GetID() == cityID {
			s.Cities[i].PurchaseFood(storeID, f, qty)
			break
		}
	}
}

func (s *State) PurchaseBeverage(
	cityID string, storeID string, b item.BeverageItem, qty int,
) {
	for i := 0; i < len(s.Cities); i++ {
		if s.Cities[i].GetID() == cityID {
			s.Cities[i].PurchaseBeverage(storeID, b, qty)
			break
		}
	}
}
