package controller

import (
	Functions "HangmanWeb/src/function"
	"fmt"
	"net/http"
	"os"
	"text/template"
)

func Treatment(w http.ResponseWriter, r *http.Request) {
	var msg string
	fmt.Println("AttemptsLeft: ", Functions.AttemptsLeft)
	if Functions.IsLetterInGuessed(r.FormValue("letter")) {
		msg = "Cette lettre a déjà été utilisée."
	} else if Functions.IsWordInGuessed(r.FormValue("letter")) {
		msg = "Ce mot a déjà été utilisée."
	} else {
		msg = ""
	}
	Functions.Guess(r.FormValue("letter"))
	Functions.DisplayHangman()
	if Functions.AttemptsLeft <= 0 {
		Functions.LooseTotal += 1
		Functions.IsGameOver = true
		http.Redirect(w, r, "/loose", http.StatusSeeOther)
		return
	} else if Functions.CheckWin() || Functions.IsWin {
		
		http.Redirect(w, r, "/win", http.StatusSeeOther)
		return
	} else {
		temp, tempErr := template.ParseFiles("./src/temps/display.html")
		if tempErr != nil {
			fmt.Println("Error parsing templates: display")
			os.Exit(1)
		}
		fmt.Println(Functions.AttemptsLeft)
		data := dataPage{
			Word:              Functions.Word,
			TabUnder:          Functions.TabUnder,
			LetterGuessedList: Functions.LetterGuessedList,
			WordGuessedList:   Functions.WordGuessedList,
			AttemptsLeft:      Functions.AttemptsLeft,
			HangmanImage:      Functions.HangmanImage,
			Message:           msg,
		}
		temp.ExecuteTemplate(w, "display", data)
	}
}
