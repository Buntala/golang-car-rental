package entity

type BookingType struct {
	BookingTypeID int    `json:"booking_type_id"  binding:"numeric" gorm:"primary_key"`
	BookingType   string `json:"booking_type" `
	Description   string `json:"description" `
}

func (BookingType) TableName() string {
	return "booking_type"
}