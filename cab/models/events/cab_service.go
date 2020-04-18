package events

import (
	cbConfig "cab/config"
	"time"
)

//ViewAvailableCabs view for cabs with latest location which are available for booking
type ViewAvailableCabs struct {
	CabID     string    `sql:"cab_id"`
	Latitude  float64   `sql:"latitude"`
	Longitude float64   `sql:"longitude"`
	Uts       time.Time `sql:"uts"`
}

//GetAvailableCabs returns all cabs which are available for booking
func GetAvailableCabs(c cbConfig.Config) ([]ViewAvailableCabs, error) {
	var rows []ViewAvailableCabs
	err := c.PG.Model(&ViewAvailableCabs{}).Select(&rows)
	return rows, err
}
