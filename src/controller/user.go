package controller

import (
	Functions "HangmanWeb/src/function"
	"fmt"
	"net/http"
	"os"
	"text/template"
)

func User(w http.ResponseWriter, r *http.Request) {
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
		temp, tempErr := template.ParseFiles("./src/temps/user.html")
		if tempErr != nil {
			fmt.Println("Error parsing templates: User")
			os.Exit(1)
		}
		temp.ExecuteTemplate(w, "user", nil)
	}
}
