package controller

import (
	"context"
	v1 "test-gf/api/v1"
	"test-gf/internal/consts"
	"test-gf/internal/model"
	"test-gf/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
)

type cRegister struct{}

var Register = cRegister{}

func (c *cRegister) Register(ctx context.Context, req *v1.RegisterReq) (res *v1.CommonOkRes, err error) {
	err = service.User().Register(ctx, model.UserRegisterInput{
		Username: req.Username,
		Password: req.Password,
		Nickname: req.Nickname,
	})

	if err != nil {
		err = gerror.WrapCode(consts.RegisterFail, err)
		return
	}

	return
}
