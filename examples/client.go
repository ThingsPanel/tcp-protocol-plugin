package main

import (
	"fmt"
	"github.com/sllt/tp-tcp-plugin/model"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", ":7653")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	isAuthed := false

	for {

		//data := make([]byte, 0)
		//data = append(data, 'T')
		//data = append(data, 'P')
		//// 0x0: data packet, 0x1: heartbeat packet
		//data = append(data, 0x0)
		//// cmd, 0x0: auth
		//data = append(data, 0x0)
		//// payload length
		//length := make([]byte, 4)
		//accessToken := "59034d98-8739-aae1-d5f6-c1f2705d5510"
		//binary.BigEndian.PutUint32(length, uint32(len(accessToken)))
		//data = append(data, length...)
		//// payload
		//data = append(data, []byte(accessToken)...)
		if !isAuthed {
			packet := model.BuildAuthPacket("2072aae3-596c-3977-4a14-b127dcef41e0")
			data := packet.Serialize()

			fmt.Println(string(data))

			_, err := conn.Write(data)
			if err != nil {
				fmt.Println(err)
				break
			}
			isAuthed = true
			time.Sleep(time.Millisecond * 200)
			continue
		}

		publish := "{\"light\": 98, \"humidity\": 30.0}"

		packet := model.BuildPublishAttributesPacket([]byte(publish))
		data := packet.Serialize()

		_, err := conn.Write(data)
		if err != nil {
			fmt.Println(err)
			break
		}

		time.Sleep(time.Second * 2)
	}
}
