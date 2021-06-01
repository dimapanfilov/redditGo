package main

import (
	"log"
	"net/http"

	"github.com/dimapanfilov/redditGo/postgres"
	"github.com/dimapanfilov/redditGo/web"
)

/*Server*/

/* Opens URL 3000 for activity */
func main() {
	store, err := postgres.NewStore("postgres://postgres:secret@localhost/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	h := web.NewHandler(store)
	http.ListenAndServe(":3000", h)
}
