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
	NbrDeaths int 
	DeathCount int
}
func Loose(w http.ResponseWriter, r *http.Request) {
	temp, tempErr := template.ParseFiles("./src/temps/loose.html")
	if tempErr != nil {
		fmt.Println("Error parsing templates: loose")
		os.Exit(1)
	}
	nbrDeaths := Functions.RandomDeath()
	DeathCount = DeathCount + nbrDeaths

	data := dataLoosePage{
		NbrDeaths: nbrDeaths,
		DeathCount: DeathCount, 
	}
	fmt.Println(nbrDeaths)
	fmt.Println(DeathCount)
	temp.ExecuteTemplate(w, "loose", data)
}
