package rpc

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPhoneTokenCacheAndInvalidTokenRefresh(t *testing.T) {
	tokens, calls := 0, 0
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		if r.Method != "POST" {
			t.Error("must use POST")
		}
		if r.URL.Path == "/token" {
			var body map[string]any
			if json.NewDecoder(r.Body).Decode(&body) != nil || body["grant_type"] != "client_credential" || body["appid"] != "test-app" || body["secret"] != "test-secret" || body["force_refresh"] != (tokens > 0) {
				t.Error("invalid stable-token request")
			}
			tokens++
			fmt.Fprintf(w, `{"access_token":"token-%d","expires_in":7200}`, tokens)
			return
		}
		var body map[string]string
		if json.NewDecoder(r.Body).Decode(&body) != nil || body["code"] == "" || r.URL.Query().Get("access_token") == "" {
			t.Error("missing phone authorization")
		}
		calls++
		if calls == 2 {
			fmt.Fprint(w, `{"errcode":42001}`)
			return
		}
		fmt.Fprint(w, `{"errcode":0,"phone_info":{"purePhoneNumber":"13800000000","countryCode":"86","watermark":{"appid":"test-app"}}}`)
	}))
	defer srv.Close()
	client := NewWeChatClient("test-app", "test-secret")
	client.tokenEndpoint, client.phoneEndpoint = srv.URL+"/token", srv.URL+"/phone"
	for i := 0; i < 3; i++ {
		phone, err := client.Phone(context.Background(), fmt.Sprintf("phone-code-%d", i))
		if err != nil || phone != "+8613800000000" {
			t.Fatal("phone exchange failed")
		}
	}
	if tokens != 2 || calls != 4 {
		t.Fatal("token caching or bounded refresh failed")
	}
}

func TestPhoneRejectsUnverifiedAndRedactsFailures(t *testing.T) {
	for _, body := range []string{
		`{"errcode":40029,"errmsg":"private-phone-secret"}`,
		`{"errcode":0,"phone_info":{"purePhoneNumber":"13800000000","countryCode":"86","watermark":{"appid":"other-app"}}}`,
		`{"errcode":0,"phone_info":{"purePhoneNumber":"bad","countryCode":"86","watermark":{"appid":"test-app"}}}`,
		`{}`, `not json`,
	} {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			if r.URL.Path == "/token" {
				fmt.Fprint(w, `{"access_token":"private-token","expires_in":7200}`)
				return
			}
			fmt.Fprint(w, body)
		}))
		client := NewWeChatClient("test-app", "test-secret")
		client.tokenEndpoint, client.phoneEndpoint = srv.URL+"/token", srv.URL+"/phone"
		phone, err := client.Phone(context.Background(), "private-code")
		srv.Close()
		if err == nil || phone != "" || strings.Contains(err.Error(), "private") || strings.Contains(err.Error(), srv.URL) {
			t.Fatal("invalid phone response accepted or leaked")
		}
	}
}

func TestPhoneTokenFailureAndCancelledContext(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"errcode":40164,"errmsg":"private-secret"}`)
	}))
	defer srv.Close()
	client := NewWeChatClient("test-app", "test-secret")
	client.tokenEndpoint = srv.URL
	if _, err := client.Phone(context.Background(), "code"); err == nil || strings.Contains(err.Error(), "private") {
		t.Fatal("token failure not safely rejected")
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := client.Phone(ctx, "code"); err == nil {
		t.Fatal("cancelled request accepted")
	}
}
