package system

import (
	"udemy-sales-management/item"
	"udemy-sales-management/state"
)

type ISystem interface {
	GetStates() []state.IState
	AddState(c state.IState)
	PurchaseFood(stateID string, cityID string, storeID string, f item.FoodItem, qty int)
	PurchaseBeverage(stateID string, cityID string, storeID string, b item.BeverageItem, qty int)
}

type System struct {
	ID     string
	States []state.IState
}

func GetInstance() ISystem {
	s := System{}

	return &s
}

func (s *System) GetStates() []state.IState {
	return s.States
}

func (s *System) AddState(st state.IState) {
	s.States = append(s.States, st)
}

func (s *System) PurchaseFood(
	stateID string, cityID string, storeID string,
	f item.FoodItem, qty int,
) {
	for i := 0; i < len(s.States); i++ {
		if s.States[i].GetID() == stateID {
			s.States[i].PurchaseFood(cityID, storeID, f, qty)
			break
		}
	}
}

func (s *System) PurchaseBeverage(
	stateID string, cityID string, storeID string,
	b item.BeverageItem, qty int,
) {
	for i := 0; i < len(s.States); i++ {
		if s.States[i].GetID() == stateID {
			s.States[i].PurchaseBeverage(cityID, storeID, b, qty)
			break
		}
	}
}
