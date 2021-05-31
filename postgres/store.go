package postgres

import (
	"fmt"

	redditgo "github.com/dimapanfilov/redditGo"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewStore(dataSourceName string) (*Store, error) {
	db, err := sqlx.Open("postgres", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	return &Store{
		ThreadStore:  NewThreadStore((db)),
		PostStore:    NewPostStore(db),
		CommentStore: NewCommentStore(db),
	}, nil
}

type Store struct {
	redditgo.ThreadStore
	redditgo.PostStore
	redditgo.CommentStore
}
