package errs

import (
	"errors"
	"fmt"
	"github.com/Chriss-Leo/qudaju-mprogram/backend/common/consts/errcode"
	"runtime/debug"

	log "github.com/sirupsen/logrus"
)

type BizError struct {
	Code string
	Err  error
}

func (e *BizError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Code, e.Err)
	}
	return e.Code
}

func NewBizError(err error, code ...string) *BizError {
	var bizErr *BizError
	if err != nil && errors.As(err, &bizErr) {
		// 已经是业务error了，直接返回
		return bizErr
	}
	//_, file, line, _ := runtime.Caller(1) // 获取调用者的位置
	if err != nil {
		//log.Errorf("\n 程序抛出已知错误: %s \n 异常点位: %s:%v", err.Error(), file, line)
		log.Errorf("程序抛出错误: %v\n调用栈:\n%s", err, debug.Stack())
	}
	if len(code) > 0 {
		return &BizError{code[0], err}
	}
	return &BizError{errcode.Failed, err}
}

func ToBizError(err error) *BizError {
	var bizErr *BizError
	if errors.As(err, &bizErr) {
		return bizErr
	}
	return &BizError{errcode.SystemBusy, err}
}
