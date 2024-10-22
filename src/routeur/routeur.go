package routeur

import (
	"HangmanWeb/src/controller"
	"fmt"
	"net/http"
	"os"
)

func InitServ() {
	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/MainMenu", controller.MainMenu)




	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/src/assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))


	fmt.Println("(http://localhost:8080/) - Server started on port:8080")
	http.ListenAndServe("localhost:8080", nil)
}
