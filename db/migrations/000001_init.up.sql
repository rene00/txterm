CREATE TABLE tx (
    id INTEGER PRIMARY KEY NOT NULL,
    date_created DATETIME NOT NULL,
    date_posted DATETIME NOT NULL,
    memo TEXT NOT NULL
);

CREATE TABLE tx_import_run (
    id INTEGER PRIMARY KEY NOT NULL,
    tx_id INTEGER NOT NULL,
    import_run_id INTEGER NOT NULL,
    FOREIGN KEY (tx_id) REFERENCES tx(id),
    FOREIGN KEY (import_run_id) REFERENCES import_run(id),
    UNIQUE(tx_id, import_run_id)
);

CREATE TABLE split (
    id INTEGER PRIMARY KEY NOT NULL,
    tx_id INTEGER NOT NULL,
    account_id INTEGER NOT NULL,
    value_num INTEGER NOT NULL,
    value_denom INTEGER NOT NULL,
    FOREIGN KEY (tx_id) REFERENCES tx(id),
    FOREIGN KEY (account_id) REFERENCES account(id)
);

CREATE TABLE import_run (
    id INTEGER PRIMARY KEY NOT NULL,
    date_created DATETIME NOT NULL,
    filename TEXT NOT NULL
);

CREATE TABLE account (
    id INTEGER PRIMARY KEY NOT NULL,
    name TEXT NOT NULL,
    description TEXT,
    code TEXT,
    account_type_id INTEGER NOT NULL,
    FOREIGN KEY (account_type_id) REFERENCES account_type(id),
    UNIQUE(name, account_type_id)
);

CREATE TABLE account_type (
    id INTEGER PRIMARY KEY NOT NULL,
    name TEXT NOT NULL,
    UNIQUE(name)
);
