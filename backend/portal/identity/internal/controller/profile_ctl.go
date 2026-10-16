package controller

import (
	"github.com/Chriss-Leo/qudaju-mprogram/backend/common/consts/errcode"
	"github.com/Chriss-Leo/qudaju-mprogram/backend/common/request"
	"github.com/Chriss-Leo/qudaju-mprogram/backend/common/request/middleware"
	"github.com/Chriss-Leo/qudaju-mprogram/backend/portal/identity/internal/service"
	"github.com/gin-gonic/gin"
	"io"
	"mime"
	"net/http"
	"strconv"
)

func UpdateProfile(svc *service.ProfileSvc) gin.HandlerFunc {
	return func(c *gin.Context) {
		value, ok := c.Get(middleware.PrincipalKey)
		principal, valid := value.(middleware.Principal)
		id, err := strconv.ParseInt(principal.UserID, 10, 64)
		if !ok || !valid || err != nil || id <= 0 {
			request.Fail(c, errcode.NoAuth)
			return
		}
		mediaType, _, _ := mime.ParseMediaType(c.GetHeader("Content-Type"))
		if mediaType == "multipart/form-data" {
			if err = c.Request.ParseMultipartForm(3 << 20); err != nil {
				request.Fail(c, errcode.ParameterError)
				return
			}
			defer c.Request.MultipartForm.RemoveAll()
		} else if mediaType == "application/x-www-form-urlencoded" {
			if err = c.Request.ParseForm(); err != nil {
				request.Fail(c, errcode.ParameterError)
				return
			}
		} else {
			request.Fail(c, errcode.ParameterError)
			return
		}
		var avatar io.Reader
		if mediaType == "multipart/form-data" {
			file, header, fileErr := c.Request.FormFile("avatar")
			if fileErr == nil {
				defer file.Close()
				if header.Size > 2<<20 {
					request.Fail(c, errcode.ParameterError)
					return
				}
				avatar = file
			} else if fileErr != http.ErrMissingFile {
				request.Fail(c, errcode.ParameterError)
				return
			}
		}
		user, err := svc.Update(c.Request.Context(), id, c.Request.FormValue("nickname"), avatar)
		if err != nil {
			request.Fail(c, err)
			return
		}
		request.OK(c, user)
	}
}
func Avatar(svc *service.ProfileSvc) gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := svc.OpenAvatar(c.Param("name"))
		if err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		defer file.Close()
		stat, err := file.Stat()
		if err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		c.Header("Content-Type", "image/jpeg")
		c.Header("X-Content-Type-Options", "nosniff")
		http.ServeContent(c.Writer, c.Request, stat.Name(), stat.ModTime(), file)
	}
}
