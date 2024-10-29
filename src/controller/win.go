package controller

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
)

func Win(w http.ResponseWriter, r *http.Request) {
	temp, tempErr := template.ParseFiles("./src/temps/Win.html")
	if tempErr != nil {
		fmt.Println("Error parsing templates: Win")
		os.Exit(1)
	}
	temp.ExecuteTemplate(w, "win", nil)
}
