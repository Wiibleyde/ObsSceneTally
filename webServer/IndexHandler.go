package webServer

import (
	"logger"
	"net/http"
	"text/template"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	logger.TrafficLogger.Println("IndexHandler called")

	t, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		logger.ErrorLogger.Println("Error parsing template")
		return
	}
	t.Execute(w, nil)
}
