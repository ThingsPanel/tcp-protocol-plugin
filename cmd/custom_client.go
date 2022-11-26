package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	connOld, err := net.Dial("tcp", ":7654")
	if err != nil {
		panic(err)
	}

	conn := connOld.(*net.TCPConn)
	conn.SetNoDelay(true)
	defer conn.Close()

	isAuthed := false
	for {
		if !isAuthed {
			data := []byte("2072aae3-596c-3977-4a14-b127dcef41e0")

			_, err := conn.Write(data)
			if err != nil {
				fmt.Println(err)
				break
			}

			buf := make([]byte, 10)
			n, _ := conn.Read(buf)
			if string(buf[:n]) == "ok" {
				isAuthed = true
			}
			continue
		}

		publish := []byte("{\"light\": 98, \"humidity\": 30.0}")

		_, err := conn.Write(publish)
		if err != nil {
			fmt.Println(err)
			break
		}
		buf := make([]byte, 100)
		n, _ := conn.Read(buf)
		fmt.Println(string(buf[:n]))

		time.Sleep(time.Second * 2)
	}
}
