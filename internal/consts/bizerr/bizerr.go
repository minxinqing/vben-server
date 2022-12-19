package bizerr

import (
	"github.com/gogf/gf/v2/errors/gerror"
)

// type BizErr struct {
// 	ge error
// }

func (biz *BizErr) Error() string {
	return (*gerror.Error)(biz).Error()
}

// func New(text string) error {
// 	biz := BizErr{}
// 	biz.ge = gerror.New(text)
// 	return biz
// }

// func Newf(format string, args ...interface{}) error {
// 	biz := BizErr{}
// 	biz.ge = gerror.Newf(format, args...)
// 	return biz
// }

type BizErr gerror.Error

func New(text string) error {
	err := gerror.New(text)
	ge, _ := err.(*gerror.Error)
	return (*BizErr)(ge)
}

func Newf(format string, args ...interface{}) error {
	err := gerror.Newf(format, args...)
	ge, _ := err.(*gerror.Error)
	return (*BizErr)(ge)
}
