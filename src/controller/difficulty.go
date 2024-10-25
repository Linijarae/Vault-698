package controller

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
)

func Difficulty(w http.ResponseWriter, r *http.Request) {
	temp, tempErr := template.ParseFiles("./src/temps/difficulty.html")
	if tempErr != nil {
		fmt.Println("Error parsing templates: Difficulty")
		os.Exit(1)
	}
	temp.ExecuteTemplate(w, "difficulty", nil)
}
