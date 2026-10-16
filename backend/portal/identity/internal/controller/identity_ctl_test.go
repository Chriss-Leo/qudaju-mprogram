package controller

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Chriss-Leo/qudaju-mprogram/backend/common/api"
	"github.com/Chriss-Leo/qudaju-mprogram/backend/portal/identity/internal/service"
	"github.com/gin-gonic/gin"
)

func TestLoginRejectsCodeOnlyBeforeDependencies(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/v1"+api.WeChatLogin, WeChatLogin(&service.IdentitySvc{}))
	for _, body := range []string{`{"code":"login-code"}`, `{"phoneCode":"phone-code"}`, `{"code":"","phoneCode":"phone-code"}`} {
		req := httptest.NewRequest("POST", "/v1"+api.WeChatLogin, strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		router.ServeHTTP(response, req)
		if !strings.Contains(response.Body.String(), `"code":"3100001"`) {
			t.Fatal("unverified registration request accepted")
		}
	}
}
