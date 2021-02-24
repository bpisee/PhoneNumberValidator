package webService

import (
	"phoneNumber/logger"

	"net/http"
)

const (
	REGULAR_STR = "^(64|0064|\\+64)\\d{8,10}$"

	PHONE_HANDLER_NAME = "phoneNumber"
	URL_QUERIES_KEY    = "number"
	BASE_URI           = "/api/validate"
	RES_NAME_          = "valid"
)

func init() {
	GetRouter().
		Queries(URL_QUERIES_KEY, REGULAR_STR).
		Name(PHONE_HANDLER_NAME)

	GetRouter().Path(BASE_URI).
		HandlerFunc(GetHandler(PHONE_HANDLER_NAME)).
		Methods("get")
}

func phoneNumberHandler(res http.ResponseWriter, req *http.Request) {
	logger.Info.Println("-->URL :", req.URL)

	resJsonMap := map[string]interface{}{RES_NAME_: false}
	logger.Info.Println("-->Result for the phone number validator:", resJsonMap)

	SendResponse(res, resJsonMap)
}
