package core

import (
	"github.com/gdamore/tcell/v2"
)

type (
	Object interface {
		Awake()
		Update()
		Dispose()
	}

	GameObject interface {
		Style() tcell.Style
		HandleEvent(tcell.Event)
		Object
		Transform
	}
)
