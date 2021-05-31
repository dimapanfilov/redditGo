package postgres

import (
	"fmt"

	redditgo "github.com/dimapanfilov/redditGo"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func NewPostStore(db *sqlx.DB) *PostStore {
	return &PostStore{
		DB: db,
	}
}

type PostStore struct {
	*sqlx.DB
}

func (s *PostStore) Post(id uuid.UUID) (redditgo.Post, error) {
	var p redditgo.Post
	if err := s.Get(&p, `SELECT * FROM posts WHERE id = $1`, id); err != nil {
		return redditgo.Post{}, fmt.Errorf("error getting post: %w", err)
	}
	return p, nil
}

func (s *PostStore) PostsByThread(threadID uuid.UUID) ([]redditgo.Post, error) {
	var pp []redditgo.Post
	if err := s.Select(&pp, `SELECT * FROM posts WHERE thread_id = $1`, threadID); err != nil {
		return []redditgo.Post{}, fmt.Errorf("error getting post: %w", err)
	}
	return pp, nil
}

func (s *PostStore) CreatePost(p *redditgo.Post) error {
	if err := s.Get(p, `INSERT INTO posts VALUES ($1, $2, $3, $4, $5) RETURNING *`,
		p.ID,
		p.ThreadID,
		p.Title,
		p.Content,
		p.Votes); err != nil {
		return fmt.Errorf("error creating post: %w", err)
	}
	return nil
}

func (s *PostStore) UpdatePost(p *redditgo.Post) error {
	if err := s.Get(p, `UPDATE posts SET title = $1, description = $2 WHERE id = $3 RETURNING *`,
		p.ThreadID,
		p.Title,
		p.Content,
		p.Votes,
		p.ID); err != nil {
		return fmt.Errorf("error updating post: %w", err)
	}
	return nil
}

func (s *PostStore) DeletePost(id uuid.UUID) error {
	if _, err := s.Exec(`DELETE FROM posts WHERE id = $1`, id); err != nil {
		return fmt.Errorf("error deleting post: %w", err)
	}
	return nil
}
