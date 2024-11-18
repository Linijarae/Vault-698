package controller

import (
	Functions "HangmanWeb/src/function"
	"fmt"
	"net/http"
)

func ExitStats(w http.ResponseWriter, r *http.Request) {
	fmt.Println("IsGameOver :", Functions.IsGameOver)
	fmt.Println("IsWin :", Functions.IsWin)
	fmt.Println("AttemptsLeft :", Functions.AttemptsLeft)

	if Functions.IsGameOver {
		http.Redirect(w, r, "/planLarge", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/treatment", http.StatusSeeOther)
	}
}
