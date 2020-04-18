package events_test

import (
	nearby "cab/cab_nearby"
	cbConfig "cab/config"
	cab_service "cab/models/events"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateBooking(t *testing.T) {
	t.Parallel()
	pg := cbConfig.InitPG()
	c := cbConfig.Config{
		PG:   pg,
		Port: "8080",
	}
	defer pg.Close()
	userLoc := nearby.UserLoc{

		SrcLatitude:   18.530219,
		SrcLongitude:  73.861250,
		DestLatitude:  18.690959,
		DestLongitude: 73.776756,
	}
	//invalid Ids test
	ID := "invalid"
	cID := "invalid"
	_, err := cab_service.CreateBooking(c, ID, userLoc.SrcLatitude, userLoc.SrcLongitude, userLoc.DestLatitude, userLoc.DestLongitude, cID)
	require.Error(t, err)

	//invalid user id test
	cID = "2d2a7434-0fb0-4033-ad25-b53e8512f4c4"
	_, err = cab_service.CreateBooking(c, ID, userLoc.SrcLatitude, userLoc.SrcLongitude, userLoc.DestLatitude, userLoc.DestLongitude, cID)
	require.Error(t, err)

	//invalid user id test of valid uuid type
	ID = "2d2a7434-0fb0-4033-ad25-b53e8512f4c4"
	_, err = cab_service.CreateBooking(c, ID, userLoc.SrcLatitude, userLoc.SrcLongitude, userLoc.DestLatitude, userLoc.DestLongitude, cID)
	require.Error(t, err)

	//invalid cab id test of valid uuid type
	ID = "e12535d4-3767-49b9-889a-0675d065b3e9"
	cID = "e12535d4-3767-49b9-889a-0675d065b3e9"
	_, err = cab_service.CreateBooking(c, ID, userLoc.SrcLatitude, userLoc.SrcLongitude, userLoc.DestLatitude, userLoc.DestLongitude, cID)
	require.Error(t, err)

	//success test for valid IDs
	ID = "e12535d4-3767-49b9-889a-0675d065b3e9"
	cID = "2d2a7434-0fb0-4033-ad25-b53e8512f4c4"
	_, err = cab_service.CreateBooking(c, ID, userLoc.SrcLatitude, userLoc.SrcLongitude, userLoc.DestLatitude, userLoc.DestLongitude, cID)
	require.NoError(t, err)
}

func TestGetBookingByUserID(t *testing.T) {
	t.Parallel()
	pg := cbConfig.InitPG()
	c := cbConfig.Config{
		PG:   pg,
		Port: "8080",
	}
	defer pg.Close()

	ID := "invalid"
	_, err := cab_service.GetBookingByUserID(c, ID)
	require.Error(t, err)

	ID = "2d2a7434-0fb0-4033-ad25-b53e8512f4c4"
	_, err = cab_service.GetBookingByUserID(c, ID)
	require.Error(t, err)

	ID = "921ed0c7-3a58-4af4-90fd-e3cc2cb76099"
	_, err = cab_service.GetBookingByUserID(c, ID)
	require.Error(t, err)

	ID = "e12535d4-3767-49b9-889a-0675d065b3e9"
	_, err = cab_service.GetBookingByUserID(c, ID)
	require.NoError(t, err)

}
