package component

import (
	"github.com/baili2023/nano/component/gamepb"
	"github.com/baili2023/nano/session"
)

type GameComponent interface {
	Component
	// 开始游戏
	RPCBegin([]*session.Session, *gamepb.Begin) error
	// 断线重连
	RPCReJoin(*session.Session) error
	// 玩家挂后台
	RPCPause(s *session.Session, _ []byte) error
	// 玩家从后台切回来
	RPCResume(s *session.Session, _ []byte) error
	// 玩家在游戏内重新连接
	// RPCReConnect(s *session.Session, _ []byte) error
	// 玩家离线
	// RPCDisconnect(s *session.Session, _ []byte) error
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
