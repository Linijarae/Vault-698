package controller

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
)

func PlanLarge(w http.ResponseWriter, r *http.Request) {
	temp, tempErr := template.ParseFiles("./src/temps/planLarge.html")
	if tempErr != nil {
		fmt.Println("Error parsing templates: PlanLarge")
		os.Exit(1)
	}
	temp.ExecuteTemplate(w, "planLarge", nil)
}
