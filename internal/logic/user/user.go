package user

import (
	"context"
	"test-gf/internal/consts/bizerr"
	"test-gf/internal/dao"
	"test-gf/internal/model"
	"test-gf/internal/model/entity"
	"test-gf/internal/service"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
)

type sUser struct{}

func init() {
	service.RegisterUser(New())
}

func New() *sUser {
	return &sUser{}
}

func (s *sUser) Register(ctx context.Context, in model.UserRegisterInput) error {
	return dao.User.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		var user *entity.User
		if err := gconv.Struct(in, &user); err != nil {
			return err
		}

		n, err := dao.User.Ctx(ctx).Where(dao.User.Columns().Username, user.Username).
			WhereOr(dao.User.Columns().Nickname, user.Nickname).Count()
		if err != nil {
			return err
		}

		if n > 0 {
			return bizerr.Newf(`账号"%s"或昵称"%s"已被占用`, user.Username, user.Nickname)
		}

		user.Salt = grand.Letters(5)
		user.Password = s.EncryptPassword(in.Password, user.Salt)
		user.Status = 1

		_, err = dao.User.Ctx(ctx).Data(user).OmitEmpty().Save()
		return err
	})
}

func (s *sUser) EncryptPassword(password, salt string) string {
	return gmd5.MustEncrypt(password + salt)
}

func (s *sUser) Login(ctx context.Context, in model.UserLoginInput) (*entity.User, error) {
	var user *entity.User
	err := dao.User.Ctx(ctx).Where(dao.User.Columns().Username, in.Username).Scan(&user)
	if err != nil {
		return nil, err
	}

	if user == nil || user.Password != s.EncryptPassword(in.Password, user.Salt) {
		return nil, bizerr.Newf("用户名或密码错误")
	}

	return user, nil
}

func (s *sUser) UserInfo(ctx context.Context, uid int) (*entity.User, error) {
	var user *entity.User
	err := dao.User.Ctx(ctx).WherePri(uid).Scan(&user)

	if err != nil {
		return nil, err
	}

	return user, nil
}
