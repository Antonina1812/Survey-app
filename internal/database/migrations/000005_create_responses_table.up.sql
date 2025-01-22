        CREATE TABLE responses (
            id SERIAL PRIMARY KEY,
            poll_id INTEGER NOT NULL,
            user_id INTEGER NOT NULL,
            created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
            FOREIGN KEY (poll_id) REFERENCES polls(id),
            FOREIGN KEY (user_id) REFERENCES users(id)
        );