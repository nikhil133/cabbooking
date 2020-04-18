package cabnearby

import (
	cbConfig "cab/config"
	cab_service "cab/models/events"
	"errors"
	"log"
	"math"

	pg "gopkg.in/pg.v4"
)

const (
	//NearestCabThreshold radius of nearby cab in km
	NearestCabThreshold = 5.0
)

//UserLoc structure to hold user location
type UserLoc struct {
	SrcLatitude   float64 `json:"source_latitude"`
	SrcLongitude  float64 `json:"source_longitude"`
	DestLatitude  float64 `json:"destination_latitude"`
	DestLongitude float64 `json:"destination_longitude"`
}

// NearestAvailableCab returns cab which is nearest within 5km
func NearestAvailableCab(c cbConfig.Config, userLoc UserLoc) (cab_service.ViewAvailableCabs, error) {

	var availableCab cab_service.ViewAvailableCabs
	if userLoc.SrcLatitude <= 0.0 || userLoc.SrcLongitude <= 0.0 || userLoc.DestLatitude <= 0.0 || userLoc.DestLongitude <= 0.0 {
		return availableCab, errors.New("User location not found")
	}

	cabs, err := cab_service.GetAvailableCabs(c)

	if err == pg.ErrNoRows {
		err = errors.New("No cabs available near by try after sometime")
		return availableCab, err
	}

	nearestCab := NearestCabThreshold
	for _, cab := range cabs {
		distance := GeoDistance(userLoc.SrcLatitude, userLoc.SrcLongitude, cab.Latitude, cab.Longitude)
		if distance < NearestCabThreshold {
			if nearestCab > distance {
				nearestCab = distance
				availableCab = cab
			}

		}
	}
	if availableCab == (cab_service.ViewAvailableCabs{}) { // ==
		err = errors.New("No cabs available near by try after sometime")
		log.Println("Error ", err)
		return availableCab, err
	}
	return availableCab, nil
}

//GeoDistance calculate distance using Haversine formula in km
func GeoDistance(lat1 float64, lon1 float64, lat2 float64, lon2 float64) float64 {
	p := 0.017453292519943295 // PI / 180
	c := math.Cos
	a := 0.5 - c((lat2-lat1)*p)/2 +
		c(lat1*p)*c(lat2*p)*
			(1-c((lon2-lon1)*p))/2

	return 12742 * math.Asin(math.Sqrt(a)) // 2 * R=12742 R = 6371 km;
}
