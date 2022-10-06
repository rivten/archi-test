DROP TABLE IF EXISTS referral;
DROP TABLE IF EXISTS ad;
DROP TABLE IF EXISTS member;

CREATE TABLE member (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    password TEXT NOT NULL
);

CREATE TABLE ad (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    owner_id INTEGER,
    price REAL,
    description TEXT,

    FOREIGN KEY(owner_id) REFERENCES member(id)
);

CREATE TABLE referral (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    a_id INTEGER,
    b_id INTEGER,
    ignored INTEGER NOT NULL, 

    FOREIGN KEY(a_id) REFERENCES member(id),
    FOREIGN KEY(b_id) REFERENCES member(id)
);
