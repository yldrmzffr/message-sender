CREATE TABLE messages (
    id SERIAL PRIMARY KEY,
    recipient VARCHAR(15) NOT NULL CHECK (length(recipient) BETWEEN 10 AND 15),
    content VARCHAR(140) NOT NULL CHECK (length(content) BETWEEN 5 AND 140),
    status VARCHAR(10) DEFAULT 'pending' NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    completed_at TIMESTAMPTZ DEFAULT NULL
);