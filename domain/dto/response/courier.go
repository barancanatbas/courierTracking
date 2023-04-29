package response

import (
	"github.com/barancanatbas/courierTracking/domain/dto"
	"time"
)

type Courier struct {
	ID          uint            `json:"id"`
	CreatedAt   time.Time       `json:"createdAt"`
	Name        string          `json:"name"`
	Phone       string          `json:"phone"`
	VehicleType dto.VehicleType `json:"vehicleType"`
}
