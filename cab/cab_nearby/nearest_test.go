package cabnearby_test

import (
	nearby "cab/cab_nearby"
	cbConfig "cab/config"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNearestAvailableCab(t *testing.T) {
	t.Parallel()
	pg := cbConfig.InitPG()
	c := cbConfig.Config{
		PG:   pg,
		Port: "8080",
	}
	defer pg.Close()

	userLoc := []nearby.UserLoc{
		nearby.UserLoc{
			SrcLatitude:   0.0,
			SrcLongitude:  0.0,
			DestLatitude:  0.0,
			DestLongitude: 0.0,
		},
		nearby.UserLoc{
			SrcLatitude:   1.0,
			SrcLongitude:  0.0,
			DestLatitude:  1.0,
			DestLongitude: 0.0,
		},
		nearby.UserLoc{
			SrcLatitude:   1.0,
			SrcLongitude:  1.0,
			DestLatitude:  1.0,
			DestLongitude: 1.0,
		},
		nearby.UserLoc{
			SrcLatitude:   18.530219,
			SrcLongitude:  73.861250,
			DestLatitude:  18.690959,
			DestLongitude: 73.776756,
		},
	}
	//test for lattitude and longitude if not grater than 0
	_, err := nearby.NearestAvailableCab(c, userLoc[0])
	require.Error(t, err)
	_, err = nearby.NearestAvailableCab(c, userLoc[1])
	require.Error(t, err)
	//test for no cars available if car not available should return error with no car available...
	_, err = nearby.NearestAvailableCab(c, userLoc[2])
	require.Error(t, err)
	//success test
	_, err = nearby.NearestAvailableCab(c, userLoc[3])
	require.NoError(t, err)
}
