package middlewares

import (
	"log"
	"net/http"
	"runtime/debug"
)

// Recovery type logs all panics that hit it to provided logger
type Recovery struct {
	Logger *log.Logger
}

// WithRecovery logs panic error along with the call stack. Tries to return 500 to client
func (m *Recovery) WithRecovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				m.Logger.Println(string(debug.Stack()))
				m.Logger.Println(err)
				http.Error(w, "Something went wrong", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
