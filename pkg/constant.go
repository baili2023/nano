package pkg

type SessionKey string

const (
	// 公共错误路由
	ONERR = "onErr"
	// 当前玩家
	CUR_PLAYER SessionKey = "player"
	// 当前牌桌
	CUR_DESK SessionKey = "desk"
	// 本地调度
	LOCAL_SCHEDULER SessionKey = "localScheduler"
)
