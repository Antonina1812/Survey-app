        CREATE TABLE questions (
            id SERIAL PRIMARY KEY,
            poll_id INTEGER NOT NULL,
            text TEXT NOT NULL,
            question_type VARCHAR(50) NOT NULL,
            created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
            FOREIGN KEY (poll_id) REFERENCES polls(id)
        );