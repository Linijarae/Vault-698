package routeur

import (
	"HangmanWeb/src/controller"
	"fmt"
	"net/http"
	"os"
)

func InitServ() {
	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/planLarge", controller.PlanLarge)
	http.HandleFunc("/difficulty", controller.Difficulty)
	http.HandleFunc("/stats", controller.Stats)
	http.HandleFunc("/exitStats", controller.ExitStats)
	http.HandleFunc("/MainMenu", controller.MainMenu)
	http.HandleFunc("/play", controller.Display)
	http.HandleFunc("/startGame", controller.StartGame)
	http.HandleFunc("/win", controller.Win)
	http.HandleFunc("/treatment", controller.Treatment)
	http.HandleFunc("/loose", controller.Loose)
	http.HandleFunc("/cheater", controller.Cheater)
	http.HandleFunc("/user", controller.User)

	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/src/assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))

	fmt.Println("(http://localhost:8080/) - Server started on port:8080")
	http.ListenAndServe("localhost:8080", nil)
}
