package webServer

import (
	"fmt"
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
	t, err := template.ParseFiles("./templates/show_camera.html")
	if err != nil {
		// Handle the error (e.g., log it, show an error page, etc.)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Println("Error parsing template:", err)
		return
	}
	t.Execute(w, htmlData)
}
