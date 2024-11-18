package controller

import (
	Functions "HangmanWeb/src/function"
	"fmt"
	"net/http"
	"os"
	"text/template"
)

type Usernames struct {
	Username string
}

var Username string

func UserTreatment(w http.ResponseWriter, r *http.Request) {
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
		Username = r.FormValue("username")
		fmt.Println(Username)
		temp, tempErr := template.ParseFiles("./src/temps/planLarge.html")
		if tempErr != nil {
			fmt.Println("Error parsing templates: planLarge")
			os.Exit(1)
		}
		data := Usernames{
			Username: Username,
		}
		temp.ExecuteTemplate(w, "planLarge", data)
	}
}
