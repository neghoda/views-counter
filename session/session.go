/*Package session allows to track unique users by setting cookie with unique id*/
package session

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"net/http"
	"net/url"
	"time"
)

var (
	sessionName   string        = "appsession"
	sessionMaxAge time.Duration = time.Minute
)

// StartSession start/rewrite session and returns sessionID along with error
func StartSession(w http.ResponseWriter, r *http.Request) (*http.Cookie, error) {
	sid, err := generateSessionID()
	if err != nil {
		return nil, err
	}
	cookie := http.Cookie{Name: sessionName, Value: url.QueryEscape(sid), MaxAge: int(sessionMaxAge)}
	http.SetCookie(w, &cookie)
	return &cookie, nil
}

// GetSessionID just calls Cookie for request with internal sessionName. http.ErrNoCookie if there no cookie set
func GetSessionID(w http.ResponseWriter, r *http.Request) (*http.Cookie, error) {
	return r.Cookie(sessionName)
}

func generateSessionID() (string, error) {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
