-- name: InsertBalances :batchexec
INSERT INTO balances (LastName, FirstName, City, Balance) VALUES ($1, $2, $3, $4);