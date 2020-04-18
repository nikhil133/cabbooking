package main

import (
	cbConfig "cab/config"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

//JSON writes json to server
func JSON(c cbConfig.Config, str string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, str)
	}
}

/*func ConnectDB()cbConfig.Config{

}*/
func main() {

	pg := cbConfig.InitPG()
	config := cbConfig.Config{
		PG:   pg,
		Port: "8080",
	}
	defer pg.Close()
	cabMux := mux.NewRouter()
	cabMux.HandleFunc("/cab/api/bookcab", UserBookCab(config)).Methods("POST")
	cabMux.HandleFunc("/cab/api/booking/history", BookingHistory(config)).Methods("POST")
	http.Handle("/", cabMux)
	http.ListenAndServe(":"+config.Port, nil)

}
