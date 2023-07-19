package webServer

import (
	"logger"
	"net/http"
	"text/template"
)

type HtmlData struct {
	CamId int
}

func ShowCamHandler(w http.ResponseWriter, r *http.Request) {
	logger.TrafficLogger.Println("ShowCamHandler called")

	camId := r.URL.Query().Get("camId")

	camIdInt := ConvertToInt(camId)

	htmlData := HtmlData{CamId: camIdInt}
	t, err := template.ParseFiles("./templates/show_camera.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		logger.ErrorLogger.Println("Error parsing template")
		return
	}
	t.Execute(w, htmlData)
}
