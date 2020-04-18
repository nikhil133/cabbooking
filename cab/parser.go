package main

import (
	"errors"
	"net/http"
	"strings"
)

//ErrorResponse struct for error response
type ErrorResponse struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

//ParseHeaderCreds to authenticate users
func ParseHeaderCreds(r *http.Request) (string, string, string, error) {
	UsernameHeader := r.Header.Get("X-Credential-Username")
	UsernameHeaderAr := strings.Split(UsernameHeader, ".")
	if len(UsernameHeaderAr) == 3 {
		//customer type, username, passsword
		return UsernameHeaderAr[0], UsernameHeaderAr[1], UsernameHeaderAr[2], nil
	}
	return "", "", "", errors.New("Unauthorised Access")

}
