package events

import (
	cbConfig "cab/config"
	"errors"
	"time"

	pg "gopkg.in/pg.v4"
)

//Booking struct for cab booking while writng to db
type Booking struct {
	ID              string    `sql:"id"`
	UserID          string    `sql:"user_id"`
	CabID           string    `sql:"cab_id"`
	SourceLatitude  float64   `sql:"source_latitude"`
	SourceLongitude float64   `sql:"source_longitude"`
	DestLatitude    float64   `sql:"dest_latitude"`
	DestLongitude   float64   `sql:"dest_longitude"`
	BookingStatus   int       `sql:"booking_status"`
	Uts             time.Time `sql:"uts"`
}

//Bookings struct for bookings done while reading from db
type Bookings struct {
	ID              string    `sql:"id" json:"booking_id"`
	CabID           string    `sql:"cab_id" json:"cab_id"`
	SourceLatitude  float64   `sql:"source_latitude" json:"source_latitude"`
	SourceLongitude float64   `sql:"source_longitude" json:"source_longitude"`
	DestLatitude    float64   `sql:"dest_latitude" json:"destination_latitude"`
	DestLongitude   float64   `sql:"dest_longitude" json:"destination_longitude"`
	BookingStatus   int       `sql:"booking_status" json:"booking_status"`
	Uts             time.Time `sql:"uts"`
}

//CreateBooking creates booking
func CreateBooking(c cbConfig.Config, userID string, usrSrcLat float64, usrSrcLon float64, usrDestLat float64, ustrDestLon float64, cabID string) (Booking, error) {
	var row Booking
	row.UserID = userID
	row.CabID = cabID
	row.SourceLatitude = usrSrcLat
	row.SourceLongitude = usrSrcLon
	row.DestLatitude = usrDestLat
	row.DestLongitude = ustrDestLon
	row.BookingStatus = 10
	row.Uts = time.Now()
	err := c.PG.Create(&row)
	return row, err
}

//GetBookingByUserID returns booking history of user by user_id
func GetBookingByUserID(c cbConfig.Config, userID string) ([]Bookings, error) {
	var rows []Bookings
	err := c.PG.Model(&Bookings{}).
		Order("uts desc").
		Where("user_id=?", userID).
		Select(&rows)
	if err == pg.ErrNoRows || len(rows) == 0 {
		err = errors.New("No Bookings available")
	}
	return rows, err
}
