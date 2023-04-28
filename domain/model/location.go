package model

import "time"

type Location struct {
	ID         int64
	Latitude   float64 // Enlem koordinatı
	Longitude  float64 // Boylam koordinatı
	RecordedAt time.Time
}
