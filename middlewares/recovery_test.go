package middlewares

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	panicMessage = "FAIL"
	panicFunc    = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { panic(panicMessage) })
)

func TestWithRecovery(t *testing.T) {
	var buf bytes.Buffer
	logger := log.New(&buf, "TEST:", log.Ldate|log.Ltime|log.Lshortfile)
	rm := &Recovery{logger}

	ts := httptest.NewServer(rm.WithRecovery(panicFunc))
	defer ts.Close()
	resp, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}

	t.Run("Recovery middlewere sets response status to 500 when posible", func(t *testing.T) {
		if s := resp.StatusCode; s != http.StatusInternalServerError {
			t.Errorf("Expected %v Status Code, got - %v", http.StatusInternalServerError, s)
		}
	})

	t.Run("Recovery middlewere writes stacktrace to provided logger", func(t *testing.T) {
		if !strings.Contains(string(buf.Bytes()), "debug.Stack") {
			t.Error("Expected to have stack tracke in provided logger")
		}
	})

	t.Run("Recovery middlewere writes panic massage to provided logger", func(t *testing.T) {
		if !strings.Contains(string(buf.Bytes()), panicMessage) {
			t.Errorf("Expected to have %v in provided logger", panicMessage)
		}
	})
}
