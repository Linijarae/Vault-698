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
	IsGameOver        bool
	IsWin             bool
}

func StartGame(w http.ResponseWriter, r *http.Request) {
	temp, tempErr := template.ParseFiles("./src/temps/display.html")
	if tempErr != nil {
		fmt.Println("Error parsing templates: display")
		os.Exit(1)
	}
	Functions.ShowTextFromFile(r.FormValue("difficulty"),)
	Functions.Underscore(Functions.Word)
	Functions.NbrRandom()
	
	data := dataPage{
	Word:              Functions.Word,
    TabUnder:          Functions.TabUnder,
    LetterGuessedList: Functions.LetterGuessedList,
    WordGuessedList:   Functions.WordGuessedList,
    AttemptsLeft:      Functions.AttemptsLeft,
	HangmanImage:      "",
    IsGameOver:        false,
    IsWin:             false,
    }
	temp.ExecuteTemplate(w, "display", data)
}
