       CREATE TABLE answers (
           id SERIAL PRIMARY KEY,
           question_id INTEGER NOT NULL,
           text TEXT NOT NULL,
           created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
           FOREIGN KEY (question_id) REFERENCES questions(id)
       );