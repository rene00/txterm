CREATE TABLE tx (
    id              INTEGER PRIMARY KEY,
    date_created    DATETIME NOT NULL,
    date_posted     DATETIME NOT NULL,
    memo            TEXT NOT NULL,
    amount_num      INTEGER NOT NULL,
    amount_den      INTEGER NOT NULL,
    import_id       INTEGER NOT NULL DEFAULT 0,
    FOREIGN KEY (import_id) REFERENCES import(id)
);
