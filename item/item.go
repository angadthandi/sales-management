package item

var (
	// food
	Sandwich = FoodItem{"Sandwich"}
	Burger   = FoodItem{"Burger"}
	Pizza    = FoodItem{"Pizza"}
	Noodles  = FoodItem{"Noodles"}
)

var (
	// beverage
	Tea    = BeverageItem{"Tea"}
	Coffee = BeverageItem{"Coffee"}
	Water  = BeverageItem{"Water"}
)

type FoodItem struct {
	item string
}

func (f FoodItem) GetItem() string {
	return f.item
}

type BeverageItem struct {
	item string
}

func (b BeverageItem) GetItem() string {
	return b.item
}
