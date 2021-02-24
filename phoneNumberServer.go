package main

import (
	"net/http"

	"phoneNumber/logger"
	"phoneNumber/webService"
)

// A phone number validator server for parsing and or verifying international phone numbers
func main() {
	const SERVER_PORT string = ":8080"

	logger.Info.Println("\n  PHONE NUMBER VALIDATOR SERVER IS RUNNING... ")

	err := http.ListenAndServe(SERVER_PORT, webService.GetRouter())

	if err != nil {
		logger.Info.Println(err)
	}
}
