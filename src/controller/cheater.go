package controller

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
)

func Cheater(w http.ResponseWriter, r *http.Request) {
	temp, tempErr := template.ParseFiles("./src/temps/cheater.html")
	if tempErr != nil {
		fmt.Println("Error parsing templates: cheater")
		os.Exit(1)
	}
	data := ""
	temp.ExecuteTemplate(w, "cheater", data)
}
