package consts

import (
	"github.com/gogf/gf/v2/errors/gcode"
)

var (
	RegisterFail = gcode.New(100001, "注册失败", nil)
	LoginFail    = gcode.New(200001, "登录失败", nil)

	GetUserInfoFail = gcode.New(300001, "获取用户信息失败", nil)
)
