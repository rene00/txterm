CREATE TABLE import (
    id INTEGER PRIMARY KEY,
    date_created DATETIME NOT NULL,
    filename TEXT NOT NULL,
    balance_amount_num INTEGER,
    balance_amount_den INTEGER,
    date_as_of DATETIME
);
