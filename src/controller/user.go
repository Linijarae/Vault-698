package controller

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
)

func User(w http.ResponseWriter, r *http.Request) {
	temp, tempErr := template.ParseFiles("./src/temps/user.html")
	if tempErr != nil {
		fmt.Println("Error parsing templates: User")
		os.Exit(1)
	}
	temp.ExecuteTemplate(w, "user", nil)
}
