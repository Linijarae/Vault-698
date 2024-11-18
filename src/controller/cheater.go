package controller

import (
	Functions "HangmanWeb/src/function"
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
	if Functions.Cheat >= 3 {
		temp.ExecuteTemplate(w, "cheater", nil)
		return
	} else {
		http.Redirect(w, r, "/display", http.StatusSeeOther)
	}
}
