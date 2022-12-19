package v1

import "github.com/gogf/gf/v2/frame/g"

type RegisterReq struct {
	g.Meta   `path:"/register" method:"post" summary:"用户注册" tags:"用户"`
	Username string `json:"username" v:"required#请输入账号" dc:"账号"`
	Password string `json:"password" v:"required#请输入密码" dc:"密码"`
	Nickname string `json:"nickname" v:"required#请输入昵称" dc:"昵称"`
}
