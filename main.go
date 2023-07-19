package main

import (
	"logger"
	"net/http"
	"tallyManager"
	"webServer"
)

func main() {
	logger.InitLogger()
	logger.InfoLogger.Println("Server starting")
	tallyManager.InitTally()
	tallyManager.ChangeCam(0, 1)
	tallyManager.ChangeCam(0, 2)
	tallyManager.ChangeCam(0, 3)
	logger.InfoLogger.Println("Tally initialized")
	logger.InfoLogger.Println("Server starting")
	webServer.InitWebserver()
	logger.InfoLogger.Println("Webserver initialized")
	logger.InfoLogger.Println("Server started on port 8080 : http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
