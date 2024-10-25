package controller

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
)

func Stats(w http.ResponseWriter, r *http.Request) {
	temp, tempErr := template.ParseFiles("./src/temps/Stats.html")
	if tempErr != nil {
		fmt.Println("Error parsing templates: Stats")
		os.Exit(1)
	}
	temp.ExecuteTemplate(w, "stats", nil)
}
