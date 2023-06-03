CREATE TABLE account (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    account_type_id INTEGER NOT NULL,
    UNIQUE(name, account_type_id),
    FOREIGN KEY (account_type_id) REFERENCES account_type(id)
);
