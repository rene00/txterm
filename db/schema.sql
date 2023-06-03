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

CREATE TABLE account (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    account_type_id INTEGER NOT NULL,
    FOREIGN KEY (account_type_id) REFERENCES account_type(id),
    UNIQUE(name, account_type_id)
);

CREATE TABLE account_type (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    UNIQUE(name)
);

CREATE TABLE account_match (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    account_id INTEGER NOT NULL,
    FOREIGN KEY (account_id) REFERENCES account(id),
    UNIQUE(name)
);

CREATE TABLE account_match_filter (
    id INTEGER PRIMARY KEY,
    filter TEXT NOT NULL,
    account_match_id INTEGER NOT NULL,
    FOREIGN KEY (account_match_id) REFERENCES account_match(id),
    UNIQUE(name, account_match_id)
);
