CREATE TABLE IF NOT EXISTS books (
    id SERIAL NOT NULL,
    name VARCHAR ( 255 ) NOT NULL,
    uid VARCHAR ( 255 ) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL,
    PRIMARY KEY ( id ),
    CONSTRAINT FK_books_user_id FOREIGN KEY (uid) REFERENCES accounts(id)  ON DELETE SET NULL
);

CREATE INDEX IF NOT EXISTS idx_books_user_id
ON books (uid)