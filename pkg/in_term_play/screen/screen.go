package screen

import (
	"errors"
	"fmt"
	"os"

	"github.com/gdamore/tcell/v2"
)

type (
	gameScreen struct {
		Screen tcell.Screen
	}
)

var (
	instance        *gameScreen
	_width, _height int
)

func InitGameScreen(width, height int) error {
	if width < 1 {
		return errors.New("width must be greater then 0")
	}

	if height < 1 {
		return errors.New("height must be greater then 0")
	}

	_width = width
	_height = height

	GetGameScreen()

	return nil
}

func GetGameScreen() *gameScreen {

	if instance == nil {
		newInstance := new(gameScreen)

		// scene.gameEventChan = gameEventChan

		screen, err := tcell.NewScreen()

		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return nil
		}
		if err = screen.Init(); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return nil
		}

		screen.SetStyle(tcell.StyleDefault.
			Foreground(tcell.ColorWhite).
			Background(tcell.ColorBlack))

		// screen.SetSize(_width, _height)
		screen.Clear()

		newInstance.Screen = screen
		instance = newInstance

		return instance
	} else {
		return instance
	}
}
