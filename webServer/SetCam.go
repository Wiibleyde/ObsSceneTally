package webServer

import (
	"logger"
	"net/http"
	"tallyManager"
)

func SetCamHandler(w http.ResponseWriter, r *http.Request) {
	logger.TrafficLogger.Println("SetCamHandler called")

	camId := r.URL.Query().Get("camId")
	layerId := r.URL.Query().Get("layerId")

	camIdInt := ConvertToInt(camId)
	layerIdInt := ConvertToInt(layerId)

	tallyManager.ChangeCam(camIdInt, layerIdInt)
	tallyManager.RemoveCamIdOfAllLayer(camIdInt, layerIdInt)
}
