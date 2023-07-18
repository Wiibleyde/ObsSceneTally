package webServer

import (
	"net/http"
	"text/template"
)

type HtmlData struct {
	CamId int
}

func ShowCamHandler(w http.ResponseWriter, r *http.Request) {
	// get camId and layerId from url
	camId := r.URL.Query().Get("camId")

	// Convert to integer
	camIdInt := ConvertToInt(camId)
	
	// create htmlData struct
	htmlData := HtmlData{CamId: camIdInt}
	t, _ := template.ParseFiles("./templates/show_camera.html")
	t.Execute(w, htmlData)
}