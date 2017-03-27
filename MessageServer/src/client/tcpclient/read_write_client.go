package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	//"strconv"
	"sync"
)

var host = flag.String("host", "localhost", "host")
var port = flag.String("port", "3563", "port")
var sendData = []byte(`{
		"S2C_Message": {
			"UserName":"Sugar",
			"Message":"Helleeeeee"
		}
	}`)

func main() {
	flag.Parse()
	conn, err := net.Dial("tcp", *host+":"+*port)

	if err != nil {
		fmt.Println("Error connecting:", err)
		os.Exit(1)
	}
	//defer conn.Close()

	fmt.Println("Connecting to " + *host + ":" + *port)

	var wg sync.WaitGroup
	wg.Add(2)

	go handleWrite(conn, &wg)
	go handleRead(conn, &wg)

	wg.Wait()
}

func handleWrite(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()

	//for i := 10; i > 0; i-- {
	//_, e := conn.Write(sendData)
	//
	//if e != nil {
	//	fmt.Println("Error to send message because of ", e.Error())
	//	//break
	//}
	//}
}

func handleRead(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	reader := bufio.NewReader(conn)
	for {
		line, err := reader.ReadString(byte('\n'))
		if err != nil {
			fmt.Print("Error to read message because of ", err)
			return
		}
		fmt.Print(line)
	}
}