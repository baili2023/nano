//go:build benchmark
// +build benchmark

package io

import (
	"log"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"testing"
	"time"

	"fmt"

	"github.com/baili2023/nano/benchmark/testdata"
	"github.com/baili2023/nano/component"
	"github.com/baili2023/nano/serialize/protobuf"
	"github.com/baili2023/nano/session"

	"github.com/baili2023/nano"
)

const (
	addr = "192.168.1.63:33251" // local address
	conc = 1000                 // concurrent client count
)

type TestHandler struct {
	component.Base
	metrics int32
	group   *nano.Group
}

func (h *TestHandler) AfterInit() {
	ticker := time.NewTicker(time.Second)

	// metrics output ticker
	go func() {
		for range ticker.C {
			println("QPS", atomic.LoadInt32(&h.metrics))
			atomic.StoreInt32(&h.metrics, 0)
		}
	}()
}

func NewTestHandler() *TestHandler {
	return &TestHandler{
		group: nano.NewGroup("handler"),
	}
}

var m = make(map[int]int, 0)

func (h *TestHandler) Ping(s *session.Session, data *testdata.Ping) error {
	atomic.AddInt32(&h.metrics, 1)
	m[1] = 1
	fmt.Print(m[1])
	return s.Push("pong", &testdata.Pong{Content: data.Content})
}

func server() {
	components := &component.Components{}
	components.Register(NewTestHandler())

	nano.Listen(addr,
		nano.WithDebugMode(),
		nano.WithSerializer(protobuf.NewSerializer()),
		nano.WithComponents(components),
	)
}

func client() {
	c := NewConnector()

	chReady := make(chan struct{})
	c.OnConnected(func() {
		chReady <- struct{}{}
	})

	if err := c.Start(addr); err != nil {
		panic(err)
	}

	c.On("HandleGameEvent", func(data interface{}) {})

	<-chReady
	for /*i := 0; i < 1; i++*/ {

		err := c.Request("Hall.Ready", nil, func(data interface{}) {
			fmt.Println(string(data.([]byte)))
			//onResult <- string(data.([]byte))
			//chWait <- struct{}{}
		})
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(1000 * time.Millisecond)
	}
}

func TestIO(t *testing.T) {
	//go server()

	// wait server startup
	time.Sleep(1 * time.Second)
	for i := 0; i < conc; i++ {
		go client()
	}

	log.SetFlags(log.LstdFlags | log.Llongfile)

	sg := make(chan os.Signal)
	signal.Notify(sg, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL)

	<-sg

	t.Log("exit")
}
