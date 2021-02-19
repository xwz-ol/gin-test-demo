package model

type Message struct {
	Username string
	Message  string
}

type UserInfo struct {
	Username string
}

type Datas struct {
	Messages []Message
	Users    []UserInfo
}


