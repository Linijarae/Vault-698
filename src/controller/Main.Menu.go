package controller

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
)

func MainMenu(w http.ResponseWriter, r *http.Request) {
	temp, tempErr := template.ParseFiles("./src/temps/MainMenu.html")
	if tempErr != nil {
		fmt.Println("Error parsing templates: MainMenu")
		os.Exit(1)
	}
	temp.ExecuteTemplate(w, "MainMennu", nil)
}
