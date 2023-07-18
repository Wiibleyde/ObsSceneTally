package webServer

import (
	"log"
	"net/http"
	"strconv"
)

// InitWebserver() is called from main.go
// It starts the webserver
func InitWebserver() {
	// set up routes
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/cam",ShowCamHandler)
	http.HandleFunc("/api/setCam", SetCamHandler)
	http.HandleFunc("/api/getLayer", GetLayerHandler)
}

func ConvertToInt(strToInvert string) int {
	// convert string to int
	convertedInt, err := strconv.Atoi(strToInvert)
	if err != nil {
		log.Fatal(err)
	}
	return convertedInt
}
