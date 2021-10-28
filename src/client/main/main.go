// @Title  聊天室客户端
// @Description
// @Author  haipinHu  08/10/2021 08:23
// @Update  haipinHu  08/10/2021 08:23
// https://www.cnblogs.com/failymao/p/15064059.html
package main

import (
	"ARQ/src/common"
	"fmt"
	"net"
	"sync"
	"time"
)

var (
	transfer *common.Transfer
	wg       = sync.WaitGroup{}
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8889")
	defer conn.Close()
	if err != nil {
		fmt.Println("连接服务端失败..")
		panic(err)
	}
	wg.Add(4)
	transfer = &common.Transfer{Conn: conn}
	ackChan := make(chan *common.Message) // 用来等待协程结束
	go transmit(ackChan, "0")
	wg.Wait()
}

// transmit 组装并发送报文段,
func transmit(ackChan chan *common.Message, expiredSec string) {

	// 发送帧
	frame := &common.Message{Sec: expiredSec, Data: "无比特错的数据"}

	_ = transfer.WritePkg(frame)
	fmt.Printf("发送%v帧;\n", expiredSec)

	// 接收ack
	go checkTimeout(ackChan, frame)
	go receive(ackChan, expiredSec)
}

// checkTimeout
func checkTimeout(ackChan chan *common.Message, frame *common.Message) {
	// 1500 or 2500
	span := time.Millisecond * 1500
	timer := time.NewTimer(span)
	for {
		select {
		case <-ackChan:
			timer.Stop() // 接收到数据后，要停止计时器
			ackChan <- &common.Message{}
			fmt.Println("       ...timer stop")
			return
		case <-timer.C: //超时判断
			_ = transfer.WritePkg(frame)
			fmt.Printf("超时重发%v帧;\n", frame.Sec)
			timer.Reset(span)
		}
	}
}

// receive ack, 并更新sec
func receive(ackChan chan *common.Message, expiredSec string) {
	for {
		ack, err := transfer.ReadPkg()
		if err != nil {
			return
		}
		if ack.Ack == expiredSec {
			// 正常接受到确认帧,
			ackChan <- &ack
			fmt.Printf("确认ack%v;\n", expiredSec)

			<-ackChan
			wg.Done()
			if expiredSec == "1" {
				expiredSec = "0"
			} else {
				expiredSec = "1"
			}
			go transmit(ackChan, expiredSec)
			return
		} else if ack.Ack != expiredSec {
			// 收到上一帧超时ACK, 丢弃
			fmt.Println("   ->收到超时ACK, 丢弃;")
		}
	}
}
