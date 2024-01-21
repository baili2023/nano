package nano

import (
	"fmt"
	"sync"

	"github.com/baili2023/nano/internal/env"
	"github.com/baili2023/nano/internal/log"
)

type Groups struct {
	mu     sync.RWMutex
	groups map[uint32]*Group
}

func NewGroups() *Groups {
	return &Groups{mu: sync.RWMutex{}, groups: map[uint32]*Group{}}
}

// Member returns specified roomNo's Group
func (cs *Groups) Member(roomNo uint32) (*Group, error) {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	g, ok := cs.groups[roomNo]
	if ok {
		return g, nil
	}

	return nil, ErrMemberNotFound
}

// Members returns all member's roomNo in current group
func (c *Groups) Members() []uint32 {
	c.mu.RLock()
	defer c.mu.RUnlock()

	var members []uint32
	for k, _ := range c.groups {
		members = append(members, k)
	}
	return members
}

// Contains check whether a roomNo is contained in current group or not
func (cs *Groups) Contains(roomNo uint32) bool {
	_, err := cs.Member(roomNo)
	return err == nil
}

// Add add group to groups
func (cs *Groups) Add(roomNo uint32, group *Group) error {

	if env.Debug {
		log.Println(fmt.Sprintf("Add group to groups , roomNo=%d", roomNo))
	}

	cs.mu.Lock()
	defer cs.mu.Unlock()

	_, ok := cs.groups[roomNo]
	if ok {
		return ErrSessionDuplication
	}

	cs.groups[roomNo] = group
	return nil
}

// Leave remove specified DeskNo related group from groups
func (c *Groups) Leave(roomNo uint32) error {

	if env.Debug {
		log.Println(fmt.Sprintf("Remove group from groups , roomNo=%d", roomNo))
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.groups, roomNo)
	return nil
}

// LeaveAll clear all group in the groups
func (c *Groups) LeaveAll() error {

	c.mu.Lock()
	defer c.mu.Unlock()
	c.groups = make(map[uint32]*Group)
	return nil
}

// Count get current member amount in the groups
func (c *Groups) Count() int {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return len(c.groups)
}
