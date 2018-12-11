package proto

type Errno int32

const (
	Errno_SUCCESS         Errno = 0
	Errno_LOGIN_FAILED    Errno = 1
	Errno_REGISTER_FAILED Errno = 2
	Errno_SEND_FAILED     Errno = 3
	Errno_SERVER_FAILED   Errno = 4
)

const (
	UserLogin        = "user_login"
	UserRegister     = "user_register"
	UserRegisterResp = "user_register_resp"
	UserLoginResp    = "user_login_resp"
	UserSendMsg      = "user_send_msg"
	UserRecvMsg      = "user_recv_msg"
	UserStatusNotify = "user_status_notify"
)

type UserStatus int32

const (
	ONLINE  UserStatus = 0
	OFFLINE UserStatus = 1
)

type ErrInfo string

func (err ErrInfo) Error() string { return "ErrorInfo: " + string(err) }

var (
	ErrUserNotExist  = ErrInfo("UserNotExist")
	ErrInvalidPasswd = ErrInfo("InvalidPasswd")
	ErrInvalidParams = ErrInfo("InvalidParams")
	ErrUserExist     = ErrInfo("ErrUserExist")
)
