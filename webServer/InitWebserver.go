package webServer

import (
	"logger"
	"net/http"
	"strconv"
)

// InitWebserver() is called from main.go
// It starts the webserver
func InitWebserver() {
	logger.TrafficLogger.Println("InitWebserver called")
	// set up routes
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/cam", ShowCamHandler)
	http.HandleFunc("/api/setCam", SetCamHandler)
	http.HandleFunc("/api/getLayer", GetLayerHandler)
}

func ConvertToInt(strToInvert string) int {
	// convert string to int
	convertedInt, err := strconv.Atoi(strToInvert)
	if err != nil {
		logger.ErrorLogger.Println("Error converting string to int")
	}
	return convertedInt
}
