package entity

import (
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	BookingID       int               `json:"booking_id" gorm:"primary_key"`
	CustomerID      int               `json:"customer_id"`
	Customer        Customer          `json:"-" gorm:"foreignKey:CustomerID"`
	CarsID          int               `json:"car_id"`
	Cars            Car               `json:"-" gorm:"foreignKey:CarsID"`
	StartTime       time.Time         `json:"start_time" binding:"omitempty"`
	EndTime         time.Time         `json:"end_time" binding:"omitempty"`
	TotalCost       int               `json:"total_cost"`
	Finished        bool              `json:"finished"`
	Discount        int               `json:"discount"`
	BookingTypeName string            `json:"booking_type"`
	BookingTypeID   int               `json:"booking_type_id"`
	BookingType     BookingType       `json:"-" gorm:"foreignKey:BookingTypeID"`
	DriverID        int               `json:"driver_id,omitempty"`
	Driver          Driver            `json:"-" gorm:"foreignKey:DriverID"`
	TotalDriverCost int               `json:"total_driver_cost"`
	DriverIncentive int               `json:"driver_incentive"`
	Deleted         gorm.DeletedAt
}

func (Booking) TableName() string {
	return "booking_table"
}