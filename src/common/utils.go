package common

import (
	"encoding/binary"
	"encoding/json"
	"net"
)

// Transfer 传输工具类, 定义了发送和读取信息的方法
type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte
}

// @title    发送信息
// @description   从Message协议数据单元读取
// @auth      hupinHu             08/9/2021 08:23
// @param     msg        参数类型         "解释"
// @return    msg        Message
func (receiver *Transfer) WritePkg(msg *Message) (err error) {
	msgStr, err := json.Marshal(*msg)
	if err != nil {
		return
	}
	binary.BigEndian.PutUint32(receiver.Buf[:4], uint32(len(msgStr)))
	n, err := receiver.Conn.Write(receiver.Buf[:4])
	if n != 4 || err != nil {
		return
	}
	n, err = receiver.Conn.Write(msgStr)
	if n != len(msgStr) || err != nil {
		return
	}
	return
}

// @title    接收信息
// @description   从套接字读取数据流, 并封装到Message协议数据单元
// @auth      hupinHu             08/9/2021 08:23
// @return    msg        Message
func (receiver *Transfer) ReadPkg() (msg Message, err error) {
	_, err = receiver.Conn.Read(receiver.Buf[:4])
	if err != nil {
		return
	}
	pkgLen := int(binary.BigEndian.Uint32(receiver.Buf[:4]))
	n, err := receiver.Conn.Read(receiver.Buf[:])
	if n != pkgLen || err != nil {
		return
	}
	err = json.Unmarshal(receiver.Buf[:pkgLen], &msg)
	return
}
