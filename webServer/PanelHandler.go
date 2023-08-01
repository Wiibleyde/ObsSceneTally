package webServer

import (
	"logger"
	"net/http"
)

func PanelHandler(w http.ResponseWriter, r *http.Request) {
	logger.TrafficLogger.Println("PanelHandler called")
	http.ServeFile(w, r, "templates/panel.html")
}