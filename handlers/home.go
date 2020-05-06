package handlers

import (
	"html/template"
	"log"
	"net/http"
	"sync"
)

// HomeHandler renders template with views counter on it
type HomeHandler struct {
	l     *log.Logger
	templ *template.Template
	once  sync.Once
}

// NewHomeHandler returns new instance of HomeHandler with specified logger
func NewHomeHandler(l *log.Logger) *HomeHandler {
	return &HomeHandler{l: l}
}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.once.Do(h.parseTemplate)
	h.templ.Execute(w, 1)
}

func (h *HomeHandler) parseTemplate() {
	h.templ = template.Must(template.ParseFiles("templates/home.html"))
}
