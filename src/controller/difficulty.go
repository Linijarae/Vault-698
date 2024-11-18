package controller

import (
	Functions "HangmanWeb/src/function"
	"fmt"
	"net/http"
	"os"
	"text/template"
)

func Difficulty(w http.ResponseWriter, r *http.Request) {
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
		temp, tempErr := template.ParseFiles("./src/temps/difficulty.html")
		if tempErr != nil {
			fmt.Println("Error parsing templates: Difficulty")
			os.Exit(1)
		}
		temp.ExecuteTemplate(w, "difficulty", nil)
	}
}
