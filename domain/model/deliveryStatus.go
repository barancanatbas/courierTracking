package model

type DeliveryStatus int

const (
	OnTheWay  DeliveryStatus = iota // 0
	Delivered                       // 1
	Cancelled                       // 2
)
