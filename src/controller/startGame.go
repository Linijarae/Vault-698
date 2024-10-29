package controller

import (
	"net/http"
)

func StartGame(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
