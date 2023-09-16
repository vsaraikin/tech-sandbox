-- name: InsertBalances :copyfrom
INSERT INTO balances (LastName, FirstName, City, Balance) VALUES ($1, $2, $3, $4);