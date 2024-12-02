package intermplay

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/gdamore/tcell/v2"
)

type (
	Game struct {
		quitTerm     chan struct{}
		quitPhys     chan struct{}
		quitc        chan struct{}
		gameover     bool
		currentScene IScene
		// manager             gameEventManager
		termEventChain      chan tcell.Event
		termEventsListeners map[TermEventsListener]bool
		wg                  sync.WaitGroup
		TermEventsListener
		sync.Mutex
	}

	TermEventsListener interface {
		HandleTermEvents(tcell.Event)
	}
)

var (
	_term            tcell.Screen
	game             *Game
	timeBeforeRender time.Time
	timeAfterRender  time.Time
)

func GetGame() *Game {
	if game == nil {
		game = new(Game)
		game.gameover = false
		game.quitTerm = make(chan struct{})
		game.quitPhys = make(chan struct{})
		game.quitc = make(chan struct{})
		// game.manager = gameEventManager{}
		game.termEventsListeners = make(map[TermEventsListener]bool)
		game.termEventChain = make(chan tcell.Event)
	}

	return game
}

func (game *Game) Close() {
	defer game.Unlock()
	game.Lock()

	close(game.quitc)
	close(game.quitTerm)
	close(game.quitPhys)
}

func (game *Game) LoadScene(scene IScene) *Game {
	defer game.Unlock()
	game.Lock()

	if game.currentScene != nil {
		game.currentScene.dispose(game.currentScene)
	}
	game.currentScene = scene

	return game
}

func (game *Game) termEventLoop() {
	defer game.wg.Done()
loop:
	for {
		select {
		case event := <-game.termEventChain:
			game.NotifyListeners(event)
		case <-game.quitTerm:
			break loop
		}
	}
}

func (game *Game) physicsLoop() {
	ticker := time.NewTicker(time.Millisecond)
	defer game.wg.Done()
	defer ticker.Stop()

loop:
	for {
		select {
		case <-game.quitPhys:
			break loop
		default:
			accumulator := defaultAccumulator
			timeBeforeRender = <-ticker.C

			for accumulator > dt {
				<-ticker.C
				game.currentScene.updatePhysics(game.currentScene)
				accumulator -= dt
			}

			alpha := accumulator / dt
			timeAfterRender = <-ticker.C
			game.currentScene.update(game.currentScene, alpha)
		}
	}
}

func (game *Game) Run() {
	game.wg.Add(2)
	go GetRenderer().ChannelEvents(game.termEventChain, game.quitc)

	game.Register(game)
	game.currentScene.awake(game.currentScene)

	go game.termEventLoop()
	go game.physicsLoop()

	game.wg.Wait()

	game.Unregister(game)
	game.currentScene.dispose(game.currentScene)

	GetRenderer().Clear()
	GetRenderer().Fini()
}

func (game *Game) handleEscape(key tcell.Key) {
	if key == tcell.KeyEscape {
		go game.Close()
	}
}

func (game *Game) HandleTermEvents(ev tcell.Event) {
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

func (game *Game) NotifyListeners(event tcell.Event) {
	for listener := range game.termEventsListeners {
		listener.HandleTermEvents(event)
	}
}

func (game *Game) Register(listener TermEventsListener) {
	defer game.Unlock()
	game.Lock()
	game.termEventsListeners[listener] = true
}

func (game *Game) Unregister(listener TermEventsListener) {
	defer game.Unlock()
	game.Lock()
	delete(game.termEventsListeners, listener)
}
