package main

import (
	Functions "HangmanWeb/src/function"
	"HangmanWeb/src/routeur"
)

func main() {
	//temps.InitTemplate()
	Functions.AttemptsLeft = 7
	Functions.IsGameOver = true
	routeur.InitServ()
	//Hangman.Menu()
}
