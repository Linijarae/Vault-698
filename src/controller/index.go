package controller

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
)

func Index(w http.ResponseWriter, r *http.Request) {

	temp, tempErr := template.ParseFiles("./src/temps/index.html")
	if tempErr != nil {
		fmt.Println("Error parsing templates: index")
		os.Exit(1)
	}
	temp.ExecuteTemplate(w, "index", nil)
}
