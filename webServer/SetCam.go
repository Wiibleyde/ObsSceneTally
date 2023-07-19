package webServer

import (
	"logger"
	"net/http"
	"tallyManager"
)

func SetCamHandler(w http.ResponseWriter, r *http.Request) {
	logger.TrafficLogger.Println("SetCamHandler called")
	// get camId and layerId from url
	camId := r.URL.Query().Get("camId")
	layerId := r.URL.Query().Get("layerId")

	// Convert to integer
	camIdInt := ConvertToInt(camId)
	layerIdInt := ConvertToInt(layerId)

	// change cam
	tallyManager.ChangeCam(camIdInt, layerIdInt)
	tallyManager.RemoveCamIdOfAllLayer(camIdInt, layerIdInt)
}
