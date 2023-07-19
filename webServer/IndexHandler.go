package webServer

import (
	"logger"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	logger.TrafficLogger.Println("IndexHandler called")
	
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("Hello World"))
}
