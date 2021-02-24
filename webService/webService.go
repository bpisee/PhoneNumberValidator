package webService

import (
	"encoding/json"
	"net/http"
	"phoneNumber/logger"

	"errors"
	"reflect"

	"github.com/gorilla/mux"
)

var webRouter *mux.Router = nil

// Build a router for handlers
func GetRouter() *mux.Router {
	if webRouter == nil {
		webRouter = mux.NewRouter()
	}
	return webRouter
}

var errorMapData = errors.New("MapData is Nil")

// Formatting the data as JSON bytes
func toJsonData(mapData interface{}) ([]byte, error) {
	if !(mapData != nil && reflect.ValueOf(mapData).Len() > 0) {
		return nil, errorMapData
	}

	resJsonData, err := json.Marshal(mapData)

	if err != nil {
		logger.Info.Println("json.Marshal failed:", err)
		return nil, err
	}

	resDataLen := len(resJsonData)
	logger.Debug.Println("<--Response length of JsonData=", resDataLen)

	return resJsonData, nil
}

type httpHandler func(http.ResponseWriter, *http.Request)

// Get one type of handler for processing the request
func GetHandler(handlerName string) httpHandler {
	logger.Info.Println("Init Handler:" + handlerName)

	switch handlerName {
	case PHONE_HANDLER_NAME:
		return phoneNumberHandler
	default:
		return defaultHandler
	}
}

// Gaven default handler
func defaultHandler(res http.ResponseWriter, req *http.Request) {
	logger.Info.Println("Default Handler Started...")

	//Do Nothing

	logger.Info.Println("Default Handler Ended")
}

// Send a  json formatted data back as response
func SendResponse(res http.ResponseWriter, mapData interface{}) {
	var lenResponse int = 0

	resData, _ := toJsonData(mapData)
	if resData != nil {
		res.Header().Set("Content-Type", "application/json")
		lenResponse, _ = res.Write(resData)
	}

	logger.Debug.Println("<--Response.Write length=", lenResponse)

}
