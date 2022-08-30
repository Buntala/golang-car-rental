package entity

type Driver struct {
	DriverID    int    `json:"driver_id"  binding:"numeric" gorm:"primary_key"`
	Name        string `json:"name" `
	Nik         string `json:"nik" binding:"omitempty,numeric"`
	PhoneNumber string `json:"phone_number" binding:"omitempty,numeric,max=12,min=5"`
	DailyCost   int    `json:"daily_cost" `
}

func (Driver) TableName() string {
	return "driver"
}
