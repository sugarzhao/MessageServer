package internal

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"server/msg"
)

func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
}

var users = make(map[gate.Agent]struct{})

func rpcNewAgent(args []interface{}) {
	log.Debug("连接之后")
	a := args[0].(gate.Agent)
	users[a] = struct{}{}
	log.Debug("called me", a.UserData())
}

func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	delete(users, a)

	userName, ok := a.UserData().(string)
	if !ok {
		return
	}

	broadcastMsg(&msg.S2C_Left{
		NumUsers: len(users),
		UserName: userName,
	}, a)
}

func broadcastMsg(msg interface{}, _a gate.Agent) {
	for a := range users {
		if a == _a {
			continue
		}
		a.WriteMsg(msg)
	}
}
