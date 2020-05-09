package handlers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/neghoda/views-couter/ctr"
	"github.com/neghoda/views-couter/session"
)

// HomeHandler renders template with views counter on it
type HomeHandler struct {
	Logger *log.Logger
	templ  *template.Template
	visits ctr.Counter
}

// NewHomeHandler returns new instance of HomeHandler with specified logger.
// Returns nil and err when failed to parse template
func NewHomeHandler(l *log.Logger) (*HomeHandler, error) {
	t, err := template.ParseFiles("templates/home.html")
	if err != nil {
		return nil, err
	}
	return &HomeHandler{l, t, 0}, nil
}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := session.GetSessionID(w, r)
	// Only track unique users within session
	if err == http.ErrNoCookie {
		_, err = session.StartSession(w, r)
		if err != nil {
			h.Logger.Printf("Failed to set cookies with err: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		h.visits.Increment()

	}
	h.templ.Execute(w, h.visits)
}

func (h *HomeHandler) parseTemplate() {
	h.templ = template.Must(template.ParseFiles("templates/home.html"))
}
