package request

import (
	"errors"
	"github.com/Chriss-Leo/qudaju-mprogram/backend/common/consts/errcode"
	"github.com/Chriss-Leo/qudaju-mprogram/backend/common/errs"
	"github.com/Chriss-Leo/qudaju-mprogram/backend/common/util"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// ResponseEntity 与 Java 结构保持一致
type ResponseEntity[T any] struct {
	Code string `json:"code"` // 错误码
	Msg  string `json:"msg"`  // 错误码说明
	Data T      `json:"data"` // 返回数据
	Time string `json:"time"` // 返回数据时间
}

func success[T any](ctx *gin.Context, data T) ResponseEntity[T] {
	return ResponseEntity[T]{
		Code: errcode.Success,
		Msg:  util.GetI18nMessage(ctx, errcode.Success),
		Data: data,
		Time: strconv.FormatInt(time.Now().Unix(), 10),
	}
}

func OK[T any](ctx *gin.Context, data T) {
	ctx.JSON(http.StatusOK, success(ctx, data))
}

func OKWithDateTime[T any](ctx *gin.Context, data T) {
	response := success(ctx, data)
	response.Time = util.FormatJavaDateTime(time.Now())
	ctx.JSON(http.StatusOK, response)
}

func OKEmpty(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, success[any](ctx, nil))
}

func OKWithCode[T any](ctx *gin.Context, code int, data T) {
	ctx.JSON(code, success(ctx, data))
}
func OKWithCodeEmpty(ctx *gin.Context, code int) {
	ctx.JSON(code, success[any](ctx, nil))
}

func Fail(ctx *gin.Context, err any) {
	var code string
	var errMsg string
	switch v := err.(type) {
	case string:
		code = v
		errMsg = util.GetI18nMessage(ctx, code)
		log.Errorf("[Request Failed] %s %s => code: %s", ctx.Request.Method, ctx.Request.URL.Path, code)
	case error:
		_ = ctx.Error(v)
		be := errs.ToBizError(v)
		code = be.Code
		errMsg = util.GetI18nMessage(ctx, code)
		if be.Err != nil {
			log.Errorf("[Request Failed] %s %s => code: %s, err: %v", ctx.Request.Method, ctx.Request.URL.Path, code, be.Err)
			if errMsg == "" || code == errcode.SystemBusy || code == errcode.Failed {
				errMsg = be.Err.Error()
			}
		}
	}
	if errMsg == "" {
		errMsg = util.GetI18nMessage(ctx, code)
	}
	// 这边需要中断执行
	ctx.AbortWithStatusJSON(http.StatusOK, ResponseEntity[any]{
		Code: code,
		Msg:  errMsg,
		Time: strconv.FormatInt(time.Now().Unix(), 10),
	})
}

// OKRemote remote的方式
func OKRemote(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusOK, data)
}

func OKRemoteString(ctx *gin.Context, data string) {
	ctx.String(http.StatusOK, data)
}

func FailRemote(ctx *gin.Context, code string) {
	errorMsg := map[string]string{
		"code": code,
	}
	ctx.JSON(999, errorMsg)
}

func FailRemoteError(ctx *gin.Context, err error) {
	var bizErr *errs.BizError
	if err != nil {
		if errors.As(err, &bizErr) {
			code := bizErr.Code
			FailRemote(ctx, code)
		} else {
			FailRemote(ctx, err.Error())
		}
	}
}
