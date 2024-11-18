package controller

import (
	Functions "HangmanWeb/src/function"
	"fmt"
	"net/http"
	"os"
	"text/template"
)

func Treatment(w http.ResponseWriter, r *http.Request) {
	temp, tempErr := template.ParseFiles("./src/temps/display.html")
	if tempErr != nil {
		fmt.Println("Error parsing templates: display")
		os.Exit(1)
	}
	Functions.Guess(r.FormValue("letter"))
	fmt.Println(Functions.AttemptsLeft)
	data := dataPage{
		Word:              Functions.Word,
		TabUnder:          Functions.TabUnder,
		LetterGuessedList: Functions.LetterGuessedList,
		WordGuessedList:   Functions.WordGuessedList,
		AttemptsLeft:      Functions.AttemptsLeft,
		HangmanImage:      Functions.HangmanImage,
	}
	temp.ExecuteTemplate(w, "display", data)
}
