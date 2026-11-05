package server

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

type pingFunc func(context.Context) error

func (f pingFunc) Ping(ctx context.Context) error { return f(ctx) }

func TestHealthDoesNotDependOnDatabase(t *testing.T) {
	handler := NewHandler(pingFunc(func(context.Context) error {
		t.Fatal("死活確認ではDBに接続しない")
		return nil
	}))
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/api/health", nil))
	if w.Code != http.StatusOK {
		t.Fatalf("status = %d", w.Code)
	}
}

func TestReadinessUnavailable(t *testing.T) {
	handler := NewHandler(pingFunc(func(ctx context.Context) error {
		if _, ok := ctx.Deadline(); !ok {
			t.Error("接続確認にタイムアウトがない")
		}
		return errors.New("private-connection-details")
	}))
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/api/ready", nil))
	if w.Code != http.StatusServiceUnavailable || strings.TrimSpace(w.Body.String()) != `{"status":"unavailable"}` {
		t.Fatalf("response = %d %s", w.Code, w.Body.String())
	}
}

func TestDatabaseReadiness(t *testing.T) {
	if os.Getenv("TEST_DATABASE") != "1" {
		t.Skip("実DBの検証は mise run test:db で実行する")
	}
	pool, err := OpenDatabase(t.Context())
	if err != nil {
		t.Fatal("DB設定を読み込めません")
	}
	defer pool.Close()
	w := httptest.NewRecorder()
	NewHandler(pool).ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/api/ready", nil))
	if w.Code != http.StatusOK || strings.TrimSpace(w.Body.String()) != `{"status":"ready"}` {
		t.Fatalf("response = %d %s", w.Code, w.Body.String())
	}
}
