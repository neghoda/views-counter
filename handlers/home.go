package handlers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/neghoda/views-couter/session"
)

// HomeHandler renders template with views counter on it
type HomeHandler struct {
	l     *log.Logger
	templ *template.Template
}

// NewHomeHandler returns new instance of HomeHandler with specified logger.
// Returns nil and err when failed to parse template
func NewHomeHandler(l *log.Logger) (*HomeHandler, error) {
	t, err := template.ParseFiles("templates/home.html")
	if err != nil {
		return nil, err
	}
	return &HomeHandler{l, t}, nil
}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := session.GetSessionID(w, r)
	if err == http.ErrNoCookie {
		_, err = session.StartSession(w, r)
		if err != nil {
			h.l.Printf("Failed to set cookies with err: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	}
	h.templ.Execute(w, 1)
}

func (h *HomeHandler) parseTemplate() {
	h.templ = template.Must(template.ParseFiles("templates/home.html"))
}
