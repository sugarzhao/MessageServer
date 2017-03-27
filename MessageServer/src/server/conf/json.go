package conf

import (
	"encoding/json"
	"io/ioutil"

	"github.com/name5566/leaf/log"
)

var Server struct {
	LogLevel   string
	LogPath    string
	WSAddr     string
	TCPAddr    string
	MaxConnNum int
}

func init() {

	/*	dirname, err := ioutil.ReadDir("../MessageServer/src/server") //获取dirname指定的目录的目录信息的有序列表。
		fmt.Println(err)
		for k, v := range dirname {
			fmt.Println(k, "=", v.Name()) //文件或目录或
			fmt.Println(v.IsDir())        //是否是目录
			fmt.Println(v.ModTime())      //文件创建时间
			fmt.Println(v.Mode())         //文件的权限
			fmt.Println(v.Size())         //文件大小
			fmt.Println(v.Sys())          //系统信息
		}*/
	data, err := ioutil.ReadFile("conf/server.json")
	if err != nil {
		log.Fatal("%v", err)
	}
	err = json.Unmarshal(data, &Server)
	if err != nil {
		log.Fatal("%v", err)
	}
}
