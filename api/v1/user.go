package v1

import "github.com/gogf/gf/v2/frame/g"

type UserInfoReq struct {
	g.Meta `path:"/user/info" method:"get" summary:"用户信息" tags:"用户"`
}

type UserInfoRes struct {
	Username string `json:"username"`
	Nickname string `json:"nickname"`
}
