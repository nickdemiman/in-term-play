package intermplay

import (
	"fmt"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
)

type (
	Game struct {
		quitq        chan struct{}
		gameover     bool
		currentScene IScene
		manager      gameEventManager
		DefaultTermEventsListener
	}
)

var (
	_term tcell.Screen
)

func NewGame() *Game {
	game := new(Game)
	game.gameover = false
	game.quitq = make(chan struct{})
	game.manager = gameEventManager{}

	return game
}

func (game *Game) Close() {
	game.currentScene.Dispose()
	GetRenderer().Fini()
	game.quitq <- struct{}{}
}

func (game *Game) LoadScene(scene IScene) {
	if game.currentScene != nil {
		// timer.GetTimer().Unregister(game.currentScene)
		game.currentScene.dispose(game.currentScene)
	}

	scene.awake(scene)
	// timer.GetTimer().Register(scene)
	game.currentScene = scene
}

func (game *Game) RunRenderer() {
	ticker := time.NewTicker(time.Millisecond * 100)

loop:
	for {
		select {
		case <-ticker.C:
			if game.currentScene != nil {
				game.currentScene.update(game.currentScene)
			}
		case <-game.quitq:
			break loop
		}
	}
}

func (game *Game) Init() {
	globalEventer = TermEventsNotifier{
		listeners: make(map[TermEventsListener]bool),
	}
	// timer.SetInterval(time.Millisecond * 100)

	go game.RunRenderer()
	go globalEventer.Run()
}

func (game *Game) Run() {
	globalEventer.Register(game)

	<-game.quitq

	globalEventer.Unregister(game)
}

func (game *Game) handleEscape(key tcell.Key) {
	if key == tcell.KeyEscape {
		game.Close()
	}
}

func (game *Game) handleTermEvents(ev tcell.Event) {

	switch ev := ev.(type) {
	case *tcell.EventKey:
		game.handleEscape(ev.Key())
	}
}

func GetRenderer() tcell.Screen {
	var err error

	if _term == nil {
		_term, err = tcell.NewScreen()

		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return nil
		}
		if err = _term.Init(); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return nil
		}

		_term.SetStyle(tcell.StyleDefault.
			Foreground(tcell.ColorWhite).
			Background(tcell.ColorBlack))

		_term.Clear()

		return _term
	} else {
		return _term
	}
}
