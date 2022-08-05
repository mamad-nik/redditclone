package redditclone

import "github.com/google/uuid"

type Thread struct {
	ID          uuid.UUID `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
}

type Post struct {
	ID       uuid.UUID `db:"id"`
	ThreadID uuid.UUID `db:"thread_id"`
	Title    string    `db:"title"`
	Content  string    `db:"content"`
	Vote     int       `db:"vote"`
}

type Comment struct {
	ID      uuid.UUID `db:"id"`
	PostID  uuid.UUID `db:"post_id"`
	Content string    `db:"content"`
	Votes   int       `db:"votes"`
}
type ThreadStore interface {
	Thread(id uuid.UUID) (Thread, error)
	Threads() ([]Thread, error)
	createThread(t *Thread) error
	updateThread(t *Thread) error
	deleteThread(id uuid.UUID) error
}
type PostStore interface {
	Post(id uuid.UUID) (Post, error)
	PostsbyThread(threadID uuid.UUID) ([]Post, error)
	createPost(t *Post) error
	updatePost(t *Post) error
	deletePost(id uuid.UUID) error
}
type CommentStore interface {
	Comment(id uuid.UUID) (Comment, error)
	CommentsByPost(postID uuid.UUID) ([]Comment, error)
	createComment(t *Comment) error
	updateComment(t *Comment) error
	deleteComment(id uuid.UUID) error
}

type Store interface {
	ThreadStore
	PostStore
	CommentStore
}
