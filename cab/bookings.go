package main

import (
	cbConfig "cab/config"
	cab_service "cab/models/events"
	auth "cab/models/verification"
	"encoding/json"
	"net/http"
)

//BookingResponse booking response body structure
type BookingResponse struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    []cab_service.Bookings `json:"data"`
}

//BookingHistory returns bookin history to client
func BookingHistory(c cbConfig.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, userName, password, err := ParseHeaderCreds(r)
		var resp BookingResponse
		var user auth.User
		resp.Status = http.StatusOK
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
		}

		booking, err := cab_service.GetBookingByUserID(c, user.ID)
		if err != nil {
			resp.Message = err.Error()
			Data, _ := json.Marshal(resp)
			JSON(c, string(Data))(w, r)
			return
		}

		resp.Data = booking
		Data, _ := json.Marshal(resp)
		JSON(c, string(Data))(w, r)
		return

	}

}
