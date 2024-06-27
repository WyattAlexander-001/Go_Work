CREATE TABLE threads (
    id UUID PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT NOT NULL
);

CREATE TABLE posts (
    id UUID PRIMARY KEY,
    thread_id UUID REFERENCES threads(id) ON DELETE CASCADE, --foreign key that ref primary key of thread id
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    votes INT DEFAULT 0
);

CREATE TABLE comments (
    id UUID PRIMARY KEY,
    post_id UUID NOT NULL REFERENCES posts(id) ON DELETE CASCADE, --foreign key that ref primary key of post id
    content TEXT NOT NULL,
    votes INT DEFAULT 0
);