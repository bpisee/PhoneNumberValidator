package main

import (
	"net/http"
)

// A phone number validator server for parsing and or verifying international phone numbers
func main() {
	const SERVER_ADDR string = ":8080"
	logger.Info.Println("\n  PHONE NUMBER VALIDATOR SERVER IS RUNNING... @", SERVER_ADDR)

}
