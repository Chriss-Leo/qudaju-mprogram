package middleware

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"testing"
)

func TestMissingTokenAndSpoofedInternalHeader(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	called := false
	r.GET("/private", RequireLogin(func(context.Context, string) (Principal, error) { return Principal{}, errors.New("invalid") }), func(c *gin.Context) { called = true; c.Status(200) })
	request := httptest.NewRequest("GET", "/private", nil)
	request.Header.Set("X-Internal-Call", "1")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)
	if w.Code != 401 || called {
		t.Fatal("internal header bypassed identity")
	}
}
func TestRolePermission(t *testing.T) {
	for _, role := range []string{"member", "other"} {
		r := gin.New()
		r.GET("/private", RequireLogin(func(context.Context, string) (Principal, error) {
			return Principal{UserID: "1", RoleCodes: []string{role}}, nil
		}), ApiPermission("member"), func(c *gin.Context) { c.String(200, "allowed") })
		w := httptest.NewRecorder()
		r.ServeHTTP(w, httptest.NewRequest("GET", "/private", nil))
		if role == "member" && w.Body.String() != "allowed" {
			t.Fatal("member rejected")
		}
		if role == "other" && w.Body.String() == "allowed" {
			t.Fatal("wrong role allowed")
		}
	}
}
