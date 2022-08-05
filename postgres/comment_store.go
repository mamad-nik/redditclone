package postgres

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/mamad-nik/redditclone"
)

func NewCommentStore(db *sqlx.DB) *CommentStore {
	return &CommentStore{
		DB: db,
	}
}

type CommentStore struct {
	*sqlx.DB
}

func (s *CommentStore) Comment(id uuid.UUID) (redditclone.Comment, error) {
	var c redditclone.Comment
	if err := s.Get(c, `SELECT * FROM comments WHERE id = $1`, id); err != nil {
		return redditclone.Comment{}, fmt.Errorf("an error occured during getting comment: %w", err)
	}
	return c, nil
}

func (s *CommentStore) CommentsByPost(postID uuid.UUID) ([]redditclone.Comment, error) {
	var cc []redditclone.Comment
	if err := s.Select(cc, `SELECT * FROM posts WHERE post_id = $1`, postID); err != nil {
		return []redditclone.Comment{}, fmt.Errorf("an error occured during getting comments: %w", err)
	}
	return cc, nil
}

func (s *CommentStore) createComment(t *redditclone.Comment) error {
	if err := s.Get(t, `INSERT INTO comments VALUES ($1, $2, $3, $4) RETURNING *`,
		t.ID,
		t.PostID,
		t.Content,
		t.Votes); err != nil {
		return fmt.Errorf("an error occured during creating comment: %w", err)
	}
	return nil
}

func (s *CommentStore) updateComment(t *redditclone.Comment) error {
	if err := s.Get(t, `UPDATE comments SET post_id = $1, content = $2, votes = $3 WHERE id = $4`,
		t.PostID,
		t.Content,
		t.Votes,
		t.ID); err != nil {
		return fmt.Errorf("an error occured during editing comment: %w", err)
	}
	return nil
}

func (s *CommentStore) deleteComment(id uuid.UUID) error {
	if _, err := s.Exec(`DELETE FROME comments WHERE id = $1`, id); err != nil {
		return fmt.Errorf("an error occured during deleting comment: %w", err)
	}
	return nil
}
