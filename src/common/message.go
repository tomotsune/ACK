// @Title  通用信息协议
// @Description 定义了通信双方所要用到的所有协议类型
// @Author  haipinHu  08/10/2021 08:23
// @Update  haipinHu  08/10/2021 08:23
package common

// 定义通信类型常量:
// Res  响应信息类型
// Login 登录信息类型
// NotifyUserStatus 用户状态信息类型
// SMS 短信息类型
// Message 信息对象, 定义了信息头和信息体

type Message struct {
	Sec  string      `json:"sec"`
	Ack  string      `json:"ack"`
	Data interface{} `json:"data"`
}
