package web

import (
	"net/http"
	"text/template"

	redditgo "github.com/dimapanfilov/redditGo"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

/* Redirects to threads link on localhost:3000/Threads*/
func NewHandler(store redditgo.Store) *Handler {
	h := &Handler{
		Mux:   chi.NewMux(),
		store: store,
	}
	h.Use(middleware.Logger)
	h.Route("/threads", func(r chi.Router) {
		r.Get("/", h.ThreadsList())
	})

	return h
}

type Handler struct {
	*chi.Mux

	store redditgo.Store
}

const threadsListHTML = `
<h1>Threads</h1>
<d1>
{{range .Threads}}
	<dt><strong>{{.Title}}</strong><dt>
	<dd>{{.Description}}</dd>
{{end}}
</d1>
`

func (h *Handler) ThreadsList() http.HandlerFunc {
	type data struct {
		Threads []redditgo.Thread
	}
	tmpl := template.Must(template.New("").Parse(threadsListHTML))
	return func(w http.ResponseWriter, r *http.Request) {
		tt, err := h.store.Threads()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, data{Threads: tt})

	}
}
