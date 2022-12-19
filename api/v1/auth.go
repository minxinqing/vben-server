package v1

import (
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

type LoginReq struct {
	g.Meta   `path:"/auth/login" method:"post" summary:"用户登录" tags:"用户"`
	Username string `json:"username" v:"required#请输入账号" dc:"账号"`
	Password string `json:"password" v:"required#请输入密码" dc:"密码"`
}

type JwtToken struct {
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
}

type LoginRes struct {
	JwtToken
}

type RefreshTokenReq struct {
	g.Meta `path:"/auth/refresh_token" method:"post"`
}

type RefreshTokenRes struct {
	JwtToken
}

type LogoutReq struct {
	g.Meta `path:"/auth/logout" method:"post"`
}
