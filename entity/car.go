package entity

type Car struct {
	CarsID         int `gorm:"primary_key"`
	Name           string
	RentPriceDaily int
	Stock          int
}

func (Car) TableName() string {
	return "cars"
}