CREATE TABLE IF NOT EXISTS record_categories (
    id SERIAL NOT NULL,
    book_id INT NOT NULL,
    title VARCHAR ( 255 ) NOT NULL,
    description VARCHAR ( 255 ) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL,
    PRIMARY KEY ( id )
);
CREATE INDEX IF NOT EXISTS idx_record_categories_books 
ON record_categories (book_id);

CREATE TABLE IF NOT EXISTS records (
    id SERIAL NOT NULL,
    book_id INT NOT NULL,
    category_id INT NOT NULL,
    amount INT NOT NULL,
    is_expense boolean DEFAULT FALSE NOT NULL,
    title VARCHAR ( 255 ) NOT NULL,
    description VARCHAR ( 255 ) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL,
    PRIMARY KEY ( id )
);
CREATE INDEX IF NOT EXISTS idx_records_categories 
ON records (category_id);
CREATE INDEX IF NOT EXISTS idx_records_books 
ON records (book_id);