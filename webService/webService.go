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

func GetRouter() *mux.Router {
	if webRouter == nil {
		webRouter = mux.NewRouter()
	}
	return webRouter
}

func SendResponse(res http.ResponseWriter, mapData interface{}) {
	var lenResponse int = 0

	resData, _ := toJsonData(mapData)
	if resData != nil {
		res.Header().Set("Content-Type", "application/json")
		lenResponse, _ = res.Write(resData)
	}

	logger.Debug.Println("<--Response.Write length=", lenResponse)

}

var errorMapData = errors.New("MapData is Nil")

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
