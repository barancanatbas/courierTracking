package model

type Delivery struct {
	ID              int64
	CustomerName    string
	CustomerPhone   string
	PickupAddress   string
	DropOffAddress  string
	PackageWeight   float64
	DeliveryFee     float64
	Status          DeliveryStatus
	CourierID       uint
	CurrentLocation []Location // geçtiği lokasyonlar.
}
