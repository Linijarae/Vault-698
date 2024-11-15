package controller

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
)

func Display(w http.ResponseWriter, r *http.Request) {
	temp, tempErr := template.ParseFiles("./src/temps/Display.html")
	if tempErr != nil {
		fmt.Println("Error parsing templates: display")
		os.Exit(1)
	}
	data := r.FormValue("difficulty")
	temp.ExecuteTemplate(w, "treatment", data)
}
