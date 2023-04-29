package model

import "gorm.io/gorm"

type Courier struct {
	gorm.Model
	Name        string
	Phone       string
	VehicleType VehicleType
	Deliveries  []Delivery // Kuryenin teslimatlarının listesi
}
