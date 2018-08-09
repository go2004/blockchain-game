package main


import (
	//"encoding/binary"
	"net"
	//"fmt"
	"github.com/golang/protobuf/proto"
	"server/msg"
	//"log"
	"encoding/binary"
	"fmt"
	//"github.com/name5566/leaf/network/protobuf"
	_ "github.com/name5566/leaf/network/protobuf"

)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:3563")
	if err != nil {
		panic(err)
	}
/*
	// Hello 消息（JSON 格式）
	// 对应游戏服务器 Hello 消息结构体
	data := []byte(`{
		"Hello": {
			"Name": "炒在工"
		}
	}`)

	// len + data
	m := make([]byte, 2+len(data))

	// 默认使用大端序
	binary.BigEndian.PutUint16(m, uint16(len(data)))

	copy(m[2:], data)

	// 发送消息
	conn.Write(m)

	m1 := make([] byte, 1024)
	len, err := conn.Read(m1)
	fmt.Print(len,string(m1))
	*/

	//test := &msg.U2LS_Login{
	//	SdkAccount: *proto.String("SdkAccount"),
	//	SdkClientId:*proto.String("SdkClientId"),
	//	SdkAccessToken:*proto.String("SdkAccessToken"),
	//}

	//test := &msg.Login{
	//	Account: "ac23",
	//	Passward: "123",
	//}

	data := []byte(`{
		"Hello": {
			"Name": "炒在工"
		}
	}`)

	// len + data
	m := make([]byte, 2+len(data))

	// 默认使用大端序
	binary.BigEndian.PutUint16(m, uint16(len(data)))

	copy(m[2:], data)

	// 发送消息
	conn.Write(m)

	m1 := make([] byte, 1024)
	len, err := conn.Read(m1)
	fmt.Print(len,string(m1))

/*

	var Processor = protobuf.NewProcessor()
	Processor.Register(&msg.Login{})

	data, err := Processor.Marshal(&msg.Login{
		Account: "a123",
		Passward: "ac23",
	})
	if err != nil {
		println("marshaling error: ", err)
	}

	// len + id + data
	m := make([]byte, 4 + len(data[1]))

	// 默认使用大端序
	binary.BigEndian.PutUint16(m, uint16(2 + len(data[1])))

	//binary.LittleEndian.PutUint16(m[2:], 10010)
	//copy(m[2:], data[0])
	copy(m[2:], data[1])

	// 发送消息
	conn.Write(m)


	buf := make([]byte, 32)
	// 接收消息
	n, err := conn.Read(buf)
	if err != nil {
		return
	}
	println("Read=",n)

	recv := &msg.LS2U_LoginResult{}
	err = proto.Unmarshal(buf[4:n], recv)
	if err != nil {

	}
	fmt.Println("Result=",recv.GetResult())
	fmt.Println("Gametoken=",recv.GetGametoken())
*/
/*
	// 进行编码
	data, err := proto.Marshal(test)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	m := make([]byte, 4+len(data))
	binary.BigEndian.PutUint16(m, uint16(len(data)))
	copy(m[4:], data)
	// 发送消息
	conn.Write(m)

	buf := make([]byte, 1024)
	// 接收消息
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	recv := &msg.LS2U_LoginResult{}
	err = proto.Unmarshal(buf[4:n], recv)
	if err != nil {

	}
	fmt.Println(recv.GetResult())
	fmt.Println(recv.GetGametoken())
*/

}

type Hello struct {
	Name string
}