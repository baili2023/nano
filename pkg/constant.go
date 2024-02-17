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
	// 当前所在牌桌号
	CUR_DESKNO SessionKey = "desk_no"
	// 当前房间
	CUR_ROOM SessionKey = "room"
)
