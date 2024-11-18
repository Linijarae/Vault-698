package controller

import (
	Functions "HangmanWeb/src/function"
	"fmt"
	"net/http"
	"os"
	"text/template"
)

var DeathCount int

type dataLoosePage struct {
	NbrDeaths  int
	DeathCount int
}

func Loose(w http.ResponseWriter, r *http.Request) {
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
		temp, tempErr := template.ParseFiles("./src/temps/loose.html")
		if tempErr != nil {
			fmt.Println("Error parsing templates: loose")
			os.Exit(1)
		}
		Functions.VarReset()
		nbrDeaths := Functions.Random()
		DeathCount = DeathCount + nbrDeaths

		data := dataLoosePage{
			NbrDeaths:  nbrDeaths,
			DeathCount: DeathCount,
		}
		fmt.Println(nbrDeaths)
		fmt.Println(DeathCount)
		temp.ExecuteTemplate(w, "loose", data)
	}
}
