package controller

import (
	Functions "HangmanWeb/src/function"
	"fmt"
	"net/http"
	"os"
	"text/template"
)

var SavedCount int

type dataWinPage struct {
	Word       string
	NbrSaved   int
	SavedCount int
}

func Win(w http.ResponseWriter, r *http.Request) {
	if !Functions.IsGameOver {
		if Functions.Cheat >= 1 {
			Functions.Cheat = 0
			http.Redirect(w, r, "/cheater", http.StatusSeeOther)
			return
		} else {
			Functions.Cheat += 1
			http.Redirect(w, r,"/treatment", http.StatusSeeOther)
		}
	} else {
		temp, tempErr := template.ParseFiles("./src/temps/win.html")
		if tempErr != nil {
			fmt.Println("Error parsing templates: win")
			os.Exit(1)
		}
		Functions.VarReset()
		nbrSaved := Functions.Random()
		SavedCount = SavedCount + nbrSaved

		data := dataWinPage{
			Word:       Functions.Word,
			NbrSaved:   nbrSaved,
			SavedCount: SavedCount,
		}
		fmt.Println(nbrSaved)
		fmt.Println(SavedCount)
		temp.ExecuteTemplate(w, "win", data)
	}
}
