package pkg

type SessionKey string

const (
	// 当前玩家
	CUR_PLAYER SessionKey = "player"
	// 当前牌桌
	CUR_DESK SessionKey = "desk"
)
