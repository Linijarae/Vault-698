package controller

import (
	HangmanWeb "HangmanWeb/src/function"
	"fmt"
	"net/http"
	"os"
	"text/template"
)
type dataPage struct {
	Word     string
	TabUnder []string
	LetterGuessedList []string
	WordGuessedList []string
	AttemptsLeft int
	HangmanImage string
	IsGameOver bool
	IsWin bool
}
func StartGame(w http.ResponseWriter, r *http.Request) {
	temp, tempErr := template.ParseFiles("./src/temps/display.html")
	if tempErr != nil {
		fmt.Println("Error parsing templates: display")
		os.Exit(1)
	}

	data := HangmanWeb.ShowTextFromFile("src/data/wordList/wordList_Easy.txt")

	temp.ExecuteTemplate(w, "display", data)

}
