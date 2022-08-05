package postgres

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/mamad-nik/redditclone"
)

func NewPostStore(db *sqlx.DB) *PostStore {
	return &PostStore{
		DB: db,
	}
}

type PostStore struct {
	*sqlx.DB
}

func (s *PostStore) Post(id uuid.UUID) (redditclone.Post, error) {
	var t redditclone.Post
	if err := s.Get(&t, `SELECT * FROM posts WHERE id = $1`, id); err != nil {
		return redditclone.Post{}, fmt.Errorf("error occured during getting post: %w", err)
	}
	return t, nil
}

func (s *PostStore) PostsbyThread(threadID uuid.UUID) ([]redditclone.Post, error) {
	var tt []redditclone.Post
	if err := s.Select(tt, `SELECT * FROM posts`); err != nil {
		return []redditclone.Post{}, fmt.Errorf("error occured during getting posts: %w", err)
	}
	return tt, nil
}

func (s *PostStore) createPost(t *redditclone.Post) error {
	if err := s.Get(t, `INSERT INTO posts VALUES ($1, $2, $3, $4, $5) RETURNING *`,
		t.ID,
		t.ThreadID,
		t.Title,
		t.Content,
		t.Vote); err != nil {
		return fmt.Errorf("error occured during writing the post: %w", err)
	}
	return nil
}

func (s *PostStore) updatePost(t *redditclone.Post) error {
	if err := s.Get(t, `UPDATE posts SET ThreadID = $1, Title = $2, Content = $3, Vote = $4 WHERE id = $5`,
		t.ThreadID,
		t.Title,
		t.Content,
		t.Vote,
		t.ID); err != nil {
		return fmt.Errorf("error occured during writing the post: %w", err)
	}
	return nil
}

func (s *PostStore) deletePost(id uuid.UUID) error {
	if _, err := s.Exec(`DELETE FROM posts WHERE id = $1`, id); err != nil {
		return fmt.Errorf("error occured during deleting post: %w", err)
	}
	return nil
}
