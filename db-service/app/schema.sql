CREATE TABLE IF NOT EXISTS balances (
    PersonID SERIAL PRIMARY KEY,
    LastName varchar(255) NOT NULL,
    FirstName varchar(255) NOT NULL,
    City varchar(255) NOT NULL,
    Balance bigint NOT NULL
);