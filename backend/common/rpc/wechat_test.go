package rpc

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestWeChatExchange(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("grant_type") != "authorization_code" || r.URL.Query().Get("js_code") == "" {
			t.Error("wrong code exchange parameters")
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"openid":"test-openid","session_key":"test-session-key"}`))
	}))
	defer srv.Close()
	client := NewWeChatClient("test-app", "test-secret")
	client.endpoint = srv.URL
	id, err := client.Exchange(context.Background(), "test-code")
	if err != nil || id != "test-openid" {
		t.Fatal("exchange failed")
	}
}
func TestWeChatExchangePlainTextJSON(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		_, _ = w.Write([]byte(`{"openid":"test-openid","session_key":"test-session-key"}`))
	}))
	defer srv.Close()
	client := NewWeChatClient("test-app", "test-secret")
	client.endpoint = srv.URL
	id, err := client.Exchange(context.Background(), "test-code")
	if err != nil || id != "test-openid" {
		t.Fatal("plain-text JSON exchange failed")
	}
}
func TestWeChatFailureDoesNotLeak(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"errcode":40029,"errmsg":"test-secret"}`))
	}))
	defer srv.Close()
	client := NewWeChatClient("test-app", "test-secret")
	client.endpoint = srv.URL
	id, err := client.Exchange(context.Background(), "test-code")
	if err == nil || id != "" || strings.Contains(err.Error(), "test-secret") {
		t.Fatal("upstream failure leak")
	}
}
func TestWeChatRejectsInvalidIdentity(t *testing.T) {
	for _, body := range []string{
		`{"errcode":40029,"errmsg":"test-secret"}`,
		`{"openid":"test-openid"}`,
		`{"session_key":"test-session-key"}`,
		`not JSON`,
	} {
		t.Run(body, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "text/plain")
				_, _ = w.Write([]byte(body))
			}))
			defer srv.Close()
			client := NewWeChatClient("test-app", "test-secret")
			client.endpoint = srv.URL
			id, err := client.Exchange(context.Background(), "test-code")
			if id != "" || err == nil || err.Error() != "wechat exchange failed" {
				t.Fatal("invalid response must fail without exposing upstream data")
			}
		})
	}
}
