package driver

type DriverVal struct {
	DriverId    int    `json:"driver_id" validate:"required,numeric"`
	Name        string `json:"name"`
	Nik         int    `json:"nik" validate:"required,numeric"`
	PhoneNumber int    `json:"phone_number" validate:"required,numeric"`
	DailyCost   int    `json:"daily_cost" validate:"required,numeric"`
}

type Get_Rules struct {
	DriverId string `json:"driver_id" validate:"required,numeric"`
}

type Post_Rules struct {
	Name        string `json:"name" validate:"required"`
	Nik         string `json:"nik" validate:"required,numeric"`
	PhoneNumber string `json:"phone_number" validate:"required,numeric"`
	DailyCost   string `json:"daily_cost" validate:"required,numeric"`
}

type Patch_Rules struct {
	DriverId    string `json:"driver_id" validate:"required,numeric"`
	Name        string `json:"name" validate:"required"`
	Nik         string `json:"nik" validate:"required,numeric"`
	PhoneNumber string `json:"phone_number" validate:"required,numeric"`
	DailyCost   string `json:"daily_cost" validate:"required,numeric"`
}

type Delete_Rules struct {
	DriverId string `json:"driver_id" validate:"required,numeric"`
}