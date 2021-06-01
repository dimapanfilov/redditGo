package web

import (
	"net/http"
	"text/template"

	redditgo "github.com/dimapanfilov/redditGo"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
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
		r.Get("/new", h.ThreadsCreate())
		r.Post("/", h.ThreadStore())
		r.Post("/{id}/delete", h.ThreadsDelete())
	})

	return h
}

type Handler struct {
	*chi.Mux

	store redditgo.Store
}

const threadsListHTML = `
<h1>Threads</h1>
<a href="threads/new">Create thread</a>
<d1>
{{range .Threads}}
	<dt><strong>{{.Title}}</strong><dt>
	<dd>{{.Description}}</dd>
	<dd>
		<form action="/threads/{{.ID}}/delete" method="POST">
			<button type="submit">Delete</button>
		</form>
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

const threadCreateHTML = `
<h1>New Thread<h1>
<form action="/threads" method="POST">
	<table>
		<tr>
			<td>Title</td>
			<td><input type="text" name="title" /></td>
		</tr>
		<tr>
			<td>Description</td>
			<td><input type="text" name="description" /></td>
		</tr>
	</table>
	<button type="submit">Create thread</button>
`

/* Opens URL /threads/new for users to create new threads */
func (h *Handler) ThreadsCreate() http.HandlerFunc {

	tmpl := template.Must(template.New("").Parse(threadCreateHTML))
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	}
}

/* Stores the above created url into the db */
func (h *Handler) ThreadStore() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		title := r.FormValue("title")
		description := r.FormValue("description")

		if err := h.store.CreateThread(&redditgo.Thread{
			ID:          uuid.New(),
			Title:       title,
			Description: description,
		}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/threads", http.StatusFound)
	}
}

/* Delete specific thread */

func (h *Handler) ThreadsDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idstr := chi.URLParam(r, "id")

		id, err := uuid.Parse(idstr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := h.store.DeleteThread(id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/threads", http.StatusFound)
	}
}
