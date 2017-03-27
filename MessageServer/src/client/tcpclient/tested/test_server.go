package tested

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/name5566/leaf/log"
	"io"
	"net"
	"sync"
)

//var addUserData = []byte(`{
//		"C2S_AddUser": {
//			"UserName": "Sugar",
//			"Message":"Hellloo"
//		}
//	}`)

var sendData = []byte(`{
		"S2C_Message": {
			"UserName":"Sugar",
			"Message":"Helleeeeee"
		}
	}`)

type C2S_AddUser struct {
	// ID 不会导出到JSON中
	//ID int `json:"-"`
	// ServerName 的值会进行二次JSON编码
	UserName string `json:"UserName"`
	Message  string `json:"Message,string"`
	// 如果 ServerIP 为空，则不输出到JSON串中
	//ServerIP string `json:"serverIP,omitempty"`
}

func Wetest() {

	var wg sync.WaitGroup
	wg.Add(1)
	for i := 1; i > 0; i-- {
		addUserData := []byte(`{
		"C2S_AddUser": {
			"UserName": "Sugar",
			"Message":"Hellloo"
		}
	}`)

		startAddUser(addUserData, sendData)
	}
	//wg.Done()
	wg.Wait()

}

func startAddUser(ad []byte, sd []byte) {

	// Hello 消息（JSON 格式）
	// 对应游戏服务器 Hello 消息结构体
	//	data = []byte(`{
	//		"C2S_AddUser": {
	//			"Name": "leaf"
	//		}
	//	}`)

	// len + data

	go writeDate(ad)

	//writeDate(sd)
}

func writeDate(data []byte) {
	conn, err := net.Dial("tcp", "127.0.0.1:3563")
	if err != nil {
		panic(err)
	}
	m := make([]byte, 2+len(data))

	// 默认使用大端序
	binary.BigEndian.PutUint16(m, uint16(len(data)))

	copy(m[2:], data)

	// 发送消息
	conn.Write(m)
	// 读取消息保持长连接（heart beat）
	readDataFully(conn)
	log.Debug("读到了消息哦")
}

func readDataFully(conn net.Conn) ([]byte, error) {
	defer func() {
		conn.Close()
		fmt.Println("close")
	}()

	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return result.Bytes(), nil
}
