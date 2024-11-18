package controller

import (
	Functions "HangmanWeb/src/function"
	"net/http"
)

func ExitStats(w http.ResponseWriter, r *http.Request) {
	Functions.IsGameOver = false
	if Functions.IsGameOver {
		http.Redirect(w, r, "/planLarge", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/treatment", http.StatusSeeOther)
	}
}
