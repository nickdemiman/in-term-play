package main

import (
	intermplay "github.com/nickdemiman/in-term-play"
	"github.com/nickdemiman/in-term-play/demos/snake/internal"
)

func main() {
	game := intermplay.GetGame()

	scene := internal.NewMainScene(0, 0, 50, 20)

	game.
		LoadScene(scene).
		Run()

}
