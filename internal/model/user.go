package model

// UserRegisterInput 用户注册
type UserRegisterInput struct {
	Username string //账号
	Password string //密码
	Nickname string //昵称
}

type UserLoginInput struct {
	Username string //账号
	Password string //密码
}
