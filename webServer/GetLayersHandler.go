package webServer

import (
	"net/http"
	"strconv"
	"tallyManager"
	"logger"
)

func GetLayersHandler(w http.ResponseWriter, r *http.Request) {
	logger.TrafficLogger.Println("GetLayersHandler called")
	// get cam
	cam := tallyManager.GetLayersInfo()

	// Write response
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(strconv.Itoa(cam[0]) + "," + strconv.Itoa(cam[1]) + "," + strconv.Itoa(cam[2])))
}
