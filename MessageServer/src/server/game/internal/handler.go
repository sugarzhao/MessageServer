package internal

import (
	"fmt"
	"reflect"
	"time"

	"github.com/name5566/leaf/gate"
	"server/msg"
)

const maxMessages = 50

var (
	messages [maxMessages]struct {
		userName string
		message  string
	}
	messageIndex int
)

var loc = time.FixedZone("", 8*3600)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

/*
小写的init（），进行自动初始化
*/
func init() {
	handleMsg(&msg.C2S_AddUser{}, handleAddUser)
	handleMsg(&msg.C2S_Message{}, handleMessage)
	handleMsg(&msg.C2S_Action{}, handleAction)
}

func handleAddUser(args []interface{}) {
	m := args[0].(*msg.C2S_AddUser)
	a := args[1].(gate.Agent)
	fmt.Println("add user name is:", m.UserName)
	fmt.Println("add user name is:", m.Message)
	a.SetUserData(m.UserName)

	for i := 0; i < maxMessages; i++ {
		index := (messageIndex + i) % maxMessages
		pm := &messages[index]
		if pm.message == "" {
			continue
		}
		a.WriteMsg(&msg.S2C_Message{
			UserName: pm.userName,
			Message:  pm.message,
		})
	}

	a.WriteMsg(&msg.S2C_Login{
		NumUsers: len(users),
	})
	broadcastMsg(&msg.S2C_Joined{
		UserName: m.UserName,
		NumUsers: len(users),
	}, a)
}

func handleMessage(args []interface{}) {
	m := args[0].(*msg.C2S_Message)
	a := args[1].(gate.Agent)
	fmt.Println("handle message")
	userName, ok := a.UserData().(string)
	if !ok {
		return
	}

	now := time.Now().In(loc)
	message := fmt.Sprintf("@%02d:%02d %s", now.Hour(), now.Minute(), m.Message)

	pm := &messages[messageIndex]
	pm.userName = userName
	pm.message = message
	messageIndex++
	if messageIndex >= maxMessages {
		messageIndex = 0
	}

	broadcastMsg(&msg.S2C_Message{
		UserName: userName,
		Message:  message,
	}, a)
}

func handleAction(args []interface{}) {
	m := args[0].(*msg.C2S_Action)
	a := args[1].(gate.Agent)

	userName, ok := a.UserData().(string)
	if !ok {
		return
	}

	switch m.Op {
	case "typing":
		broadcastMsg(&msg.S2C_Typing{
			UserName: userName,
		}, a)
	case "stop typing":
		broadcastMsg(&msg.S2C_StopTyping{
			UserName: userName,
		}, a)
	}
}
