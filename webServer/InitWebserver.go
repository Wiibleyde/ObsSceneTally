package webServer

import (
	"logger"
	"net/http"
	"strconv"
)

func InitWebserver() {
	logger.TrafficLogger.Println("InitWebserver called")
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/cam", ShowCamHandler)
	http.HandleFunc("/panel", PanelHandler)
	http.HandleFunc("/api/reset", ResetHandler)
	http.HandleFunc("/api/setCam", SetCamHandler)
	http.HandleFunc("/api/getLayer", GetLayerHandler)
	http.HandleFunc("/api/getLayers", GetLayersHandler)
}

func ConvertToInt(strToInvert string) int {
	convertedInt, err := strconv.Atoi(strToInvert)
	if err != nil {
		logger.ErrorLogger.Println("Error converting string to int")
	}
	return convertedInt
}
