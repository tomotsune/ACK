// @Title  聊天室服务端
// @Description
// @Author  haipinHu  08/10/2021 08:23
// @Update  haipinHu  08/10/2021 08:23
package main

import (
	"ARQ/src/common"
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Println("listen 8889...")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	defer listen.Close()
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			continue
		}
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	transfer := common.Transfer{Conn: conn}
	expiredSec := "0"
	for {
		msg, err := transfer.ReadPkg()
		if err != nil {
			return
		}
		if msg.Sec == expiredSec {
			ack := &common.Message{Ack: expiredSec, Data: "无比特错的数据"}
			time.Sleep(time.Millisecond * 2000)
			transfer.WritePkg(ack)
			fmt.Printf("server: 确认%v;\n", expiredSec)
			if expiredSec == "0" {
				expiredSec = "1"
			} else {
				expiredSec = "0"
			}
		} else {
			lastSec := expiredSec
			if lastSec == "0" {
				lastSec = "1"
			} else {
				lastSec = "0"
			}
			fmt.Printf("server: 丢弃重复的%v帧, 从传确认%v;\n", lastSec, lastSec)
			ack := &common.Message{Ack: expiredSec, Data: "无比特错的数据"}
			time.Sleep(time.Millisecond * 1000)
			transfer.WritePkg(ack)
		}
	}
}
