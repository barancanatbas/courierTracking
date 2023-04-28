package model

type Courier struct {
	ID          int64
	Name        string
	Phone       string
	VehicleType VehicleType

	Deliveries []Delivery // Kuryenin teslimatlarının listesi
}
