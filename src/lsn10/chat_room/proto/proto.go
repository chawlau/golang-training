package proto

type Message struct {
	Cmd  string `json:"cmd"`
	Data string `json:"data"`
}

type LoginReq struct {
	UserId string `json:"user_id"`
	Passwd string `json:"passwd"`
}

type LoginResp struct {
	Err      Errno    `json:"err_no"`
	UserList []string `json:"user_list"`
}

type RegisterReq struct {
	User *UserInfo `json:"user_info"`
}

type RegisterResp struct {
	Err Errno `json:"err_no"`
}

type UserStatusNotifyResp struct {
	UserId string     `json:"user_id"`
	Status UserStatus `json:"status"`
}

type UserSendMessageReq struct {
	UserId string `json:"user_id"`
	Data   string `json:"data"`
}

type UserRecvMessageResp struct {
	UserId string `json:"user_id"`
	Data   string `json:"data"`
}

type UserInfo struct {
	UserId    string     `json:"user_id"`
	Passwd    string     `json:"passwd"`
	NickName  string     `json:"nick_name"`
	Sex       string     `json:"sex"`
	Header    string     `json:"header"`
	LastLogin string     `json:"last_login"`
	Status    UserStatus `json:"status"`
}
