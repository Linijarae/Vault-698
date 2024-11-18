package controller

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
	"HangmanWeb/src/function"
)



type dataStatPage struct {
	WinTotal   int
	LooseTotal int
	SavedCount int
	DeathCount int
}

func Stats(w http.ResponseWriter, r *http.Request) {
	temp, tempErr := template.ParseFiles("./src/temps/Stats.html")
	if tempErr != nil {
		fmt.Println("Error parsing templates: Stats")
		os.Exit(1)
	}

	data := dataStatPage{
		SavedCount: SavedCount,
		DeathCount: DeathCount,
		WinTotal:   Functions.WinTotal,
		LooseTotal: Functions.LooseTotal,
	}

	temp.ExecuteTemplate(w, "stats", data)
}
