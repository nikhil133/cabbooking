package main

import (
	nearby "cab/cab_nearby"
	cbConfig "cab/config"
	cab_service "cab/models/events"
	auth "cab/models/verification"
	"encoding/json"
	"net/http"
)

//Response structure for json response from server side
type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

//UserBookCab books a cab for user request
func UserBookCab(c cbConfig.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var userLoc nearby.UserLoc
		var err error
		var resp Response
		resp.Status = http.StatusOK
		err = json.NewDecoder(r.Body).Decode(&userLoc)
		if err != nil {
			resp.Status = http.StatusBadRequest
			resp.Message = err.Error()
			Data, _ := json.Marshal(resp)
			JSON(c, string(Data))(w, r)

			return
		}

		token, userName, password, err := ParseHeaderCreds(r)
		var user auth.User

		if err != nil {
			resp.Status = http.StatusBadRequest
			resp.Message = "Bad Request"
			Data, _ := json.Marshal(resp)
			JSON(c, string(Data))(w, r)
			return

		}
		err = user.UserAuthentication(c, token, userName, password)
		if err != nil {
			resp.Status = http.StatusUnauthorized
			resp.Message = err.Error()
			Data, _ := json.Marshal(resp)
			JSON(c, string(Data))(w, r)
			return
		}

		alotedCab, err := nearby.NearestAvailableCab(c, userLoc)
		if err != nil {

			resp.Message = err.Error()
			Data, _ := json.Marshal(resp)
			JSON(c, string(Data))(w, r)
			return

		}
		booked, err := cab_service.CreateBooking(c, user.ID, userLoc.SrcLatitude, userLoc.SrcLongitude, userLoc.DestLatitude, userLoc.DestLongitude, alotedCab.CabID)
		if err != nil {
			resp.Message = "Booking Failed"
			Data, _ := json.Marshal(resp)
			JSON(c, string(Data))(w, r)
			return

		}
		resp.Status = 200
		resp.Message = "Cab booked successfully alloted cab " + alotedCab.CabID + "with booking Id " + booked.ID
		Data, _ := json.Marshal(resp)
		JSON(c, string(Data))(w, r)
		return

	}
}
