CREATE TABLE IF NOT EXISTS users
(
    id SERIAL not null,
    name text,
    second_name text,
    tel text,
    email text,
    password_hash varchar(64) not null,
    balance numeric,
    PRIMARY KEY (id)
);

INSERT into users(id, name, second_name, tel, email, password_hash, balance) VALUES (1, 'Alice', 'Miller', '+112300000', 'alice.miller@mail.com', 'cb824cd5fe4950a77e36776d275f8f7039682babd490d5da3bc8fd31f4c2254c', 1000);
INSERT into users(id, name, second_name, tel, email, password_hash, balance) VALUES (2, 'Bob', 'Brown', '+113500000', 'bob.brown@mail.com', 'bc786c379d8b4334faa1f5ed4428d53ed5fbf6247a5974a72eac7fd5c13410d8', 1000);

CREATE TABLE IF NOT EXISTS sessions
(
    id SERIAL not null,
    user_id integer,
    secret text,
    expire timestamp,
    PRIMARY KEY (id)
    );

INSERT into sessions(user_id, secret, expire) VALUES (1, 'abd52071ae09da447b528cfe27fe47ae5e4788285ef04cea0359cea649fdfc0f84e1ee62368c937511164854fa53715e1f4abbc333ee7c0b042c975cc944aed6', '1759051102');
INSERT into sessions(user_id, secret, expire) VALUES (2, 'd6dc3cbea6c9ed0e8e126759f4b156516cc9af7da69ee9f8bbde53f8feafb270e88a5c7b6b553cf942187c0f1a97b83ee030cf67d96ea0d92c342abab2664e5b', '1759051102');


CREATE TABLE IF NOT EXISTS transactions
(
    id SERIAL not null,
    user_id_from integer not null  REFERENCES  users(id),
    user_id_to integer not null  REFERENCES  users(id),
    amount numeric NOT NULL,
    PRIMARY KEY (id)
);