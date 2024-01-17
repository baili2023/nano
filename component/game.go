package component

import (
	"github.com/baili2023/nano/component/gamepb"
	"github.com/baili2023/nano/session"
)

type GameComponent interface {
	Component
	Begin([]*session.Session, *gamepb.Begin) error
}

type GameBase struct{}

// Init was called to initialize the component.
func (c *GameBase) Init() {}

// AfterInit was called after the component is initialized.
func (c *GameBase) AfterInit() {}

// BeforeShutdown was called before the component to shutdown.
func (c *GameBase) BeforeShutdown() {}

// Shutdown was called to shutdown the component.
func (c *GameBase) Shutdown() {}
