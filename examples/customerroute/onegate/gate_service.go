package onegate

import (
	"fmt"
	"github.com/pingcap/errors"
	"nano/component"
	"nano/examples/cluster/protocol"
	"nano/session"
)

type RegisterService struct {
	component.Base
	nextGateUid int64
}

func newRegisterService() *RegisterService {
	return &RegisterService{}
}

type (
	RegisterRequest struct {
		Nickname string `json:"nickname"`
	}
	RegisterResponse struct {
		Code int `json:"code"`
	}
)

func (bs *RegisterService) Login(s *session.Session, msg *RegisterRequest) error {
	bs.nextGateUid++
	uid := bs.nextGateUid
	s.Bind(uid)
	fmt.Println("Login uid:", uid)
	chat := &protocol.JoinRoomRequest{
		Nickname:  msg.Nickname,
		GateUid:   uid,
		MasterUid: uid,
	}
	if err := s.RPC("ChatRoomService.JoinRoom", chat); err != nil {
		return errors.Trace(err)
	}
	return s.Response(&RegisterResponse{})
}
