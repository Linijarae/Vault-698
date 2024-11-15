package controller

import (
	Functions "HangmanWeb/src/function"
	"fmt"
	"net/http"
	"os"
	"text/template"
)

type dataPage struct {
	Word              string
	TabUnder          []string
	LetterGuessedList []string
	WordGuessedList   []string
	AttemptsLeft      int
	HangmanImage      string
}

func StartGame(w http.ResponseWriter, r *http.Request) {
	if !Functions.IsGameOver {
		Functions.VarReset()
		Functions.ShowTextFromFile(r.FormValue("difficulty"))
		Functions.Underscore(Functions.Word)
		Functions.NbrRandom()
		temp, tempErr := template.ParseFiles("./src/temps/display.html")
		if tempErr != nil {
			fmt.Println("Error parsing templates: display")
			os.Exit(1)
		}
		data := dataPage{
			Word:              Functions.Word,
			TabUnder:          Functions.TabUnder,
			LetterGuessedList: Functions.LetterGuessedList,
			WordGuessedList:   Functions.WordGuessedList,
			AttemptsLeft:      7,
			HangmanImage:      "../static/img/Erreur0.png",
		}
		temp.ExecuteTemplate(w, "display", data)
	} else {

		http.Redirect(w, r, "/gameOver", http.StatusSeeOther)
	}
}
