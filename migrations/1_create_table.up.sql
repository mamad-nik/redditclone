CREATE TABLE threads(
    id UUID PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT    NOT NULL
);
CREATE TABLE posts(
    id UUID PRIMARY KEY,
    thread_id UUID NOT NULL,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    votes INT DEFAULT 0,
    FOREIGN KEY(thread_id) references threads(id)
);
CREATE TABLE comments(
    id UUID PRIMARY KEY,
    post_id UUID NOT NULL,
    content TEXT NOT NULL,
    votes INT DEFAULT 0,
    FOREIGN KEY(post_id) references posts(id)
);