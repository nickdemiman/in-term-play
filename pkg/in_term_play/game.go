package core

import (
	"sync"
	"time"

	"github.com/NickDemiman/InTermPlay/pkg/in_term_play/screen"
	"github.com/NickDemiman/InTermPlay/pkg/in_term_play/timer"
	"github.com/gdamore/tcell/v2"
)

type (
	Game struct {
		gameEventChan chan GameEvent
		gameover      bool
		currentScene  Scene
	}

	TermEventsNotifier struct {
		listeners map[TermEventsListener]bool
		sync.Mutex
	}
	TermEventsListener interface {
		handleTermEvents(tcell.Event)
	}
)

var globalEventer TermEventsNotifier

func (t *TermEventsNotifier) Run() {
	for {
		event := screen.GetGameScreen().Screen.PollEvent()
		t.Notify(event)
	}
}

func (t *TermEventsNotifier) Notify(event tcell.Event) {
	for listener := range t.listeners {
		listener.handleTermEvents(event)
	}
}

func (t *TermEventsNotifier) Register(listener TermEventsListener) {
	defer t.Unlock()
	t.Lock()
	t.listeners[listener] = true
}

func (t *TermEventsNotifier) Unregister(listener TermEventsListener) {
	defer t.Unlock()
	t.Lock()
	delete(t.listeners, listener)
}

func NewGame() *Game {
	game := new(Game)
	game.gameover = false
	game.gameEventChan = make(chan GameEvent)

	return game
}

func (game *Game) Close() {
	game.currentScene.Dispose()
}

func (game *Game) Run() {
	defer close(game.gameEventChan)

	var (
		err   error
		scene Scene
	)

	globalEventer.listeners = make(map[TermEventsListener]bool)
	timer.SetInterval(time.Millisecond * 100)

	go timer.GetTimer().Run()
	go globalEventer.Run()

	globalEventer.Register(game)

	scene, err = NewScene(0, 0, 50, 20, &game.gameEventChan)

	if err != nil {
		panic(err)
	}

	game.currentScene = scene
	game.currentScene.Awake()
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
