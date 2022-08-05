package postgres

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/mamad-nik/redditclone"
)

func NewThreadStore(db *sqlx.DB) *ThreadStore {
	return &ThreadStore{
		DB: db,
	}
}

type ThreadStore struct {
	*sqlx.DB
}

func (s *ThreadStore) Thread(id uuid.UUID) (redditclone.Thread, error) {
	var t redditclone.Thread
	if err := s.Get(&t, `SELECT * FROM threads WHERE id = $1`, id); err != nil {
		return redditclone.Thread{}, fmt.Errorf("error occured during getting thread: %w", err)
	}
	return t, nil
}

func (s *ThreadStore) Threads() ([]redditclone.Thread, error) {
	var tt []redditclone.Thread
	if err := s.Select(tt, `SELECT * FROM threads`); err != nil {
		return []redditclone.Thread{}, fmt.Errorf("error occured during getting threads: %w", err)
	}
	return tt, nil
}

func (s *ThreadStore) createThread(t *redditclone.Thread) error {
	if err := s.Get(t, `INSERT INTO threads VALUES ($1, $2, $3) RETURNING *`,
		t.ID,
		t.Title,
		t.Description); err != nil {
		return fmt.Errorf("error occured during writing the thread: %w", err)
	}
	return nil
}

func (s *ThreadStore) updateThread(t *redditclone.Thread) error {
	if err := s.Get(t, `UPDATE threads SET title = $1, description = $2 WHERE id = $3 RETURNING *`,
		t.Title,
		t.Description,
		t.ID); err != nil {
		return fmt.Errorf("error occured during writing the thread: %w", err)
	}
	return nil
}

func (s *ThreadStore) deleteThread(id uuid.UUID) error {
	if _, err := s.Exec(`DELETE FROME threads WHERE id = $1`, id); err != nil {
		return fmt.Errorf("an error occured during deleting: %w", err)
	}
	return nil
}
