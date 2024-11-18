package controller

import (
	Functions "HangmanWeb/src/function"
	"fmt"
	"net/http"
	"os"
	"text/template"
)

func Display(w http.ResponseWriter, r *http.Request) {
	if !Functions.IsGameOver {
		if Functions.Cheat >= 1 {
			Functions.Cheat = 0
			http.Redirect(w, r, "/cheater", http.StatusSeeOther)
			return
		} else {
			Functions.Cheat += 1
			http.Redirect(w, r, "/treatment", http.StatusSeeOther)
		}
	} else {
		temp, tempErr := template.ParseFiles("./src/temps/Display.html")
		if tempErr != nil {
			fmt.Println("Error parsing templates: display")
			os.Exit(1)
		}
		data := r.FormValue("difficulty")
		temp.ExecuteTemplate(w, "treatment", data)
	}
}
