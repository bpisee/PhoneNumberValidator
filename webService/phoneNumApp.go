package webService

import (
	"phoneNumber/logger"

	"net/http"
	"regexp"

	"strings"
)

const (
	//Valid phone number patterns, it will be put into a config file.
	REGULAR_STR = "^(64|0064|\\+64|0061|61|\\+61|0086|86|\\+86)\\d{8,10}$"

	PHONE_HANDLER_NAME = "phoneNumber"
	URL_QUERIES_KEY    = "number"
	BASE_URI           = "/api/validate"
	RES_NAME_          = "valid"
)

// Add the phone number validator as a handler into the HTTP　router
func init() {
	GetRouter().
		Queries(URL_QUERIES_KEY, REGULAR_STR).
		Name(PHONE_HANDLER_NAME)

	//Accepts the HTTP [GET] method only
	GetRouter().Path(BASE_URI).
		HandlerFunc(GetHandler(PHONE_HANDLER_NAME)).
		Methods("get")
}

// A phone number validator to verify a international phone number listed in URL queries
func phoneNumberHandler(res http.ResponseWriter, req *http.Request) {
	logger.Info.Println("-->URL :", req.URL)

	phoneNumStr := req.FormValue(URL_QUERIES_KEY)

	isValid := false
	if len(phoneNumStr) > 0 {
		phoneNumStr = strings.ReplaceAll(phoneNumStr, " ", "")

		isValid = verifyPhoneNumber(phoneNumStr)
	}

	resJsonMap := map[string]interface{}{RES_NAME_: isValid}
	logger.Info.Println("-->Result from the phone number validator:", resJsonMap)

	SendResponse(res, resJsonMap)
}

var reg = regexp.MustCompile(REGULAR_STR)

// verify a international phone number
func verifyPhoneNumber(phoneNumStr string) bool {
	return reg.MatchString(phoneNumStr)
}
