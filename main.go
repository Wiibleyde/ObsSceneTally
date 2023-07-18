package main

import (
	"fmt"
	"net/http"
	"tallyManager"
	"webServer"
)

func main() {
	tallyManager.InitTally()
	tallyManager.ChangeCam(0, 1)
	tallyManager.ChangeCam(0, 2)
	tallyManager.ChangeCam(0, 3)
	webServer.InitWebserver()
	fmt.Println("Server started")
	fmt.Println("http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
