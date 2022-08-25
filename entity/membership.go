package entity

type Membership struct {
	MembershipID int    `json:"membership_id" gorm:"primary_key"`
	Name         string `json:"name"`
	Discount     int    `json:"discount" `
}