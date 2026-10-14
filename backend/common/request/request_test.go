package request_test

import (
	"encoding/json"
	"errors"
	"github.com/Chriss-Leo/qudaju-mprogram/backend/common/errs"
	"github.com/Chriss-Leo/qudaju-mprogram/backend/common/request"
	"github.com/Chriss-Leo/qudaju-mprogram/backend/common/request/middleware"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestBusinessResponseContract(t *testing.T) {
	gin.SetMode(gin.TestMode)
	tests := []struct {
		name    string
		handler gin.HandlerFunc
		code    string
		wantNil bool
	}{
		{"success", func(c *gin.Context) { request.OK(c, gin.H{"id": "123"}) }, "1000000", false},
		{"empty", func(c *gin.Context) { request.OKEmpty(c) }, "1000000", true},
		{"parameter", func(c *gin.Context) { request.Fail(c, "3100001") }, "3100001", true},
		{"business", func(c *gin.Context) { request.Fail(c, errs.NewBizError(nil, "noAuth")) }, "noAuth", true},
		{"plain-error", func(c *gin.Context) { request.Fail(c, errors.New("sample failure")) }, "3010107", true},
		{"panic", func(c *gin.Context) { panic("sample panic") }, "3000001", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gin.New()
			r.Use(middleware.UnKnowErrorHandler())
			r.GET("/test", tt.handler)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, httptest.NewRequest("GET", "/test", nil))
			if w.Code != 200 {
				t.Fatalf("HTTP status = %d", w.Code)
			}
			var body map[string]any
			if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil {
				t.Fatal(err)
			}
			if len(body) != 4 || body["code"] != tt.code {
				t.Fatalf("unexpected envelope: %#v", body)
			}
			if _, ok := body["msg"].(string); !ok {
				t.Fatal("msg must be string")
			}
			ts, ok := body["time"].(string)
			if !ok {
				t.Fatal("time must be string")
			}
			seconds, err := strconv.ParseInt(ts, 10, 64)
			if err != nil || seconds <= 0 || seconds > 9999999999 {
				t.Fatal("time must be Unix seconds")
			}
			if tt.wantNil && body["data"] != nil {
				t.Fatal("data must be null")
			}
		})
	}
}

func TestFailAbortsFollowingHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	called := false
	r.GET("/test", func(c *gin.Context) { request.Fail(c, "noAuth"); return }, func(c *gin.Context) { called = true })
	r.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest("GET", "/test", nil))
	if called {
		t.Fatal("Fail must abort following handlers")
	}
}
