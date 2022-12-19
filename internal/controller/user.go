package controller

import (
	"context"
	v1 "test-gf/api/v1"
	"test-gf/internal/consts"
	"test-gf/internal/consts/bizerr"
	"test-gf/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
)

type cUser struct{}

var User = cUser{}

func (c *cUser) Info(ctx context.Context, req *v1.UserInfoReq) (res *v1.UserInfoRes, err error) {
	uid := gconv.Int(service.Auth().GetIdentity(ctx))
	userInfo, err := service.User().UserInfo(ctx, uid)

	if err != nil {
		err = gerror.WrapCode(consts.GetUserInfoFail, err)
		return
	}

	if userInfo == nil {
		err = gerror.WrapCode(consts.GetUserInfoFail, bizerr.Newf("用户不存在"))
		return
	}

	err = gconv.Struct(userInfo, &res)
	if err != nil {
		err = gerror.WrapCode(consts.GetUserInfoFail, err)
		return
	}
	return
}
