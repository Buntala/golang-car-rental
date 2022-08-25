package entity

type Customer struct {
	CustomerID     int        `json:"customer_id" gorm:"primary_key"`
	Name           string     `json:"name"`
	Nik            string     `json:"nik" binding:"omitempty,numeric"`
	PhoneNumber    string     `json:"phone_number" binding:"omitempty,numeric"`
	MembershipID   int        `json:"membership_id,omitempty"`
	Membership     Membership `gorm:"ForeignKey:MembershipID"`
	MembershipName string     `json:"membership_name" gorm:"-"`
}

func (Customer) TableName() string {
	return "customer_gorm"
}