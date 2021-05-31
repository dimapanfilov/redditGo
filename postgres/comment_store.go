package postgres

import (
	"fmt"

	redditgo "github.com/dimapanfilov/redditGo"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type CommentStore struct {
	*sqlx.DB
}

func (s *CommentStore) Comment(id uuid.UUID) (redditgo.Comment, error) {
	var c redditgo.Comment
	if err := s.Get(&c, `SELECT * FROM posts WHERE id = $1`, id); err != nil {
		return redditgo.Comment{}, fmt.Errorf("error getting post: %w", err)
	}
	return c, nil
}

func (s *CommentStore) CommentsByPost(postID uuid.UUID) ([]redditgo.Comment, error) {
	var cc []redditgo.Comment
	if err := s.Select(&cc, `SELECT * FROM posts WHERE thread_id = $1`, postID); err != nil {
		return []redditgo.Comment{}, fmt.Errorf("error getting post: %w", err)
	}
	return cc, nil
}

func (s *CommentStore) CreateComment(c *redditgo.Comment) error {
	if err := s.Get(c, `INSERT INTO comments VALUES ($1, $2, $3, $4, $5) RETURNING *`,
		c.ID,
		c.PostID,
		c.Content,
		c.Votes); err != nil {
		return fmt.Errorf("error creating comment: %w", err)
	}
	return nil
}

func (s *CommentStore) UpdateComment(c *redditgo.Comment) error {
	if err := s.Get(c, `UPDATE comments SET post_id = $1, content = $2, votes = $3 WHERE id = $4 RETURNING *`,
		c.ID,
		c.Content,
		c.Votes,
		c.ID); err != nil {
		return fmt.Errorf("error updating comment: %w", err)
	}
	return nil
}

func (s *CommentStore) DeleteComment(id uuid.UUID) error {
	if _, err := s.Exec(`DELETE FROM comments WHERE id = $1`, id); err != nil {
		return fmt.Errorf("error deleting comment: %w", err)
	}
	return nil
}
