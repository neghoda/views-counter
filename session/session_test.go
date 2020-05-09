package session

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStartSession(t *testing.T) {
	r := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	StartSession(w, r)
	var cookie *http.Cookie
	for _, c := range w.Result().Cookies() {
		if c.Name == sessionName {
			cookie = c
		}
	}

	t.Run("StartSession sets session cookie", func(t *testing.T) {
		if cookie == nil {
			t.Errorf("Failed to find cookie with name \"%v\"", sessionName)
			t.FailNow()
		}
	})

	t.Run("StartSession sets session cookie with some value", func(t *testing.T) {
		if cookie.Value == "" {
			t.Error("Expected to cookie value to not be empty")
		}
	})
}
