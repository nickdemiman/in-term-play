package main

import (
	intermplay "github.com/nickdemiman/in-term-play"
	"github.com/nickdemiman/in-term-play/demos/snake/internal"
)

func main() {
	game := intermplay.NewGame()

	scene := internal.NewMainScene(0, 0, 50, 20)

	game.Init()
	game.LoadScene(scene)
	game.Run()

}
