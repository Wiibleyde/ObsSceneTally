package webServer

import (
	"net/http"
	"strconv"
	"tallyManager"
)

func GetLayerHandler(w http.ResponseWriter, r *http.Request) {
	// get layerId from url
	layerId := r.URL.Query().Get("layerId")

	// Convert to integer
	layerIdInt := ConvertToInt(layerId)

	// get cam
	cam := tallyManager.GetLayerInfo(layerIdInt)

	// write cam to response
	w.Write([]byte(strconv.Itoa(cam)))
}
