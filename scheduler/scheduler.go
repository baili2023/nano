// Copyright (c) nano Authors. All Rights Reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package scheduler

import (
	"fmt"
	"runtime/debug"
	"sync/atomic"
	"time"

	"github.com/baili2023/nano/internal/env"
	"github.com/baili2023/nano/internal/log"
)

const (
	messageQueueBacklog = 1 << 10
	sessionCloseBacklog = 1 << 8
)

// LocalScheduler schedules task to a customized goroutine
type LocalScheduler interface {
	Schedule(Task)
}

type Task func()

type Hook func()

// 按会话进行绑定
type QueueLocalScheduler struct {
	chTasks chan Task
}

// QueueLocalScheduler
func NewQueueLocalScheduler() *QueueLocalScheduler {
	qs := &QueueLocalScheduler{chTasks: make(chan Task, 1<<8)}
	//消费消息队列
	go qs.Sched()
	return qs
}

// Schedule 读协程往自定义消息队列写入任务
func (localScheduler *QueueLocalScheduler) Schedule(task Task) {
	select {
	case localScheduler.chTasks <- task:
	case <-time.After(6 * time.Second):
		log.Println("localScheduler.PushTask  Timeout ")
	}

}

// Sched 消息自定义协程任务
func (localScheduler *QueueLocalScheduler) Sched() {
	defer func() {
		close(localScheduler.chTasks)
	}()
	for {
		select {
		case f := <-localScheduler.chTasks:
			try(f)

		case <-chDie:
			return
		}
	}
}

// DefaultScheduler 无消息队列的任务处理
//
//	每个读线程 会同时进行访问 存在并发问题
type DefaultScheduler struct{}

func NewDefaultScheduler() *DefaultScheduler {
	return &DefaultScheduler{}
}

// Schedule 每个读线程直接执行
func (defaultScheduler *DefaultScheduler) Schedule(task Task) {
	try(task)
}

var (
	chDie   = make(chan struct{})
	chExit  = make(chan struct{})
	chTasks = make(chan Task, 1<<8)
	started int32
	closed  int32
)

func try(f func()) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(fmt.Sprintf("Handle message panic: %+v\n%s", err, debug.Stack()))
		}
	}()
	f()
}

func Sched() {
	if atomic.AddInt32(&started, 1) != 1 {
		return
	}

	ticker := time.NewTicker(env.TimerPrecision)
	defer func() {
		ticker.Stop()
		close(chExit)
	}()

	for {
		select {
		case <-ticker.C:
			cron()
		case f := <-chTasks:
			try(f)
		case <-chDie:
			return
		}
	}
}

func Close() {
	if atomic.AddInt32(&closed, 1) != 1 {
		return
	}
	close(chDie)
	<-chExit
	log.Println("Scheduler stopped")
}

func PushTask(task Task) {

	select {
	case chTasks <- task:
	case <-time.After(6 * time.Second):
		log.Println("PushTask  Timeout ")
	}
	// log.Println("push task success")
}
