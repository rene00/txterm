CREATE TABLE tx (
    id              INTEGER PRIMARY KEY,
    date_created DATETIME NOT NULL,
    date_posted     DATETIME NOT NULL,
    memo            TEXT NOT NULL,
    amount_num      INTEGER NOT NULL,
    amount_den      INTEGER NOT NULL,
    import_id       INTEGER NOT NULL DEFAULT 0,
    FOREIGN KEY (import_id) REFERENCES import(id)
);

CREATE TABLE import (
    id INTEGER PRIMARY KEY,
    date_created DATETIME NOT NULL,
    filename TEXT NOT NULL,
    balance_amount_num INTEGER,
    balance_amount_den INTEGER,
    date_as_of DATETIME
);
