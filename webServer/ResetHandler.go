package webServer

import (
	"logger"
	"net/http"
	"tallyManager"
)

func ResetHandler(w http.ResponseWriter, r *http.Request) {
	logger.TrafficLogger.Println("ResetHandler called")

	tallyManager.ResetTally()
}
