package model

// Member   用户对象，定义了用户的基础信息
type Member struct {
	Mobile string `json:"mobile"`
	Pwd    string `json:"pwd"`
}
