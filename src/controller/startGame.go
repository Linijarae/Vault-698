package controller

import (
	Functions "HangmanWeb/src/function"
	"fmt"
	"net/http"
	"os"
	"text/template"
)

type dataPage struct {
	IsGameOver        bool
	Word              string
	TabUnder          []string
	LetterGuessedList []string
	WordGuessedList   []string
	AttemptsLeft      int
	HangmanImage      string
	Message           string
}

func StartGame(w http.ResponseWriter, r *http.Request) {
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
		Functions.IsGameOver = false
		Functions.ShowTextFromFile(r.FormValue("difficulty"))
		Functions.Underscore(Functions.Word)
		Functions.NbrRandom()
		temp, tempErr := template.ParseFiles("./src/temps/display.html")
		if tempErr != nil {
			fmt.Println("Error parsing templates: display")
			os.Exit(1)
		}
		fmt.Println(Functions.AttemptsLeft)
		data := dataPage{
			IsGameOver:        Functions.IsGameOver,
			Word:              Functions.Word,
			TabUnder:          Functions.TabUnder,
			LetterGuessedList: Functions.LetterGuessedList,
			WordGuessedList:   Functions.WordGuessedList,
			AttemptsLeft:      7,
			HangmanImage:      "../static/img/Erreur0.png",
		}
		temp.ExecuteTemplate(w, "display", data)
	}

}
