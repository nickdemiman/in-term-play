package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
	_            = iota
	SecondButton = iota
	FirstButton  = iota
)

var currentSel = FirstButton

func main() {
	app := tview.NewApplication()

	styleDefault := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	styleActivated := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorGreen)

	flex := tview.NewFlex()
	flex.SetDirection(tview.FlexColumnCSS).SetBorder(true).SetTitle(" Snake By NikDem ")
	btnNewGame := tview.
		NewButton("New Game").
		SetSelectedFunc(func() {
			// app.Stop()
		}).
		SetActivatedStyle(styleActivated).
		SetStyle(styleDefault)

	btnCloseGame := tview.NewButton("Close Game").
		SetSelectedFunc(func() {
			app.Stop()
		}).
		SetActivatedStyle(styleActivated).
		SetStyle(styleDefault)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyUp:
			currentSel++
		case tcell.KeyDown:
			currentSel--
		}

		if currentSel < SecondButton {
			currentSel = SecondButton
		}

		if currentSel > FirstButton {
			currentSel = FirstButton
		}

		switch currentSel {
		case FirstButton:
			app.SetFocus(btnNewGame)
		case SecondButton:
			app.SetFocus(btnCloseGame)
		}
		return event
	})

	flex.AddItem(btnNewGame, 0, 1, false)
	flex.AddItem(btnCloseGame, 0, 1, false)

	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}
