       CREATE TABLE response_answers (
           id SERIAL PRIMARY KEY,
           response_id INTEGER NOT NULL,
           question_id INTEGER NOT NULL,
           answer_id INTEGER NOT NULL,
           FOREIGN KEY (response_id) REFERENCES responses(id),
           FOREIGN KEY (question_id) REFERENCES questions(id),
           FOREIGN KEY (answer_id) REFERENCES answers(id)
       );