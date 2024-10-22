package routeur

import (
	"HangmanWeb/src/controller"
	"fmt"
	"net/http"
)

func InitServ() {
	http.HandleFunc("/", controller.Index)

	fmt.Println("(http://localhost:8080/) - Server started on port:8080")
	http.ListenAndServe("localhost:8080", nil)
}
