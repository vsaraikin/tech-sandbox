-- SET TRANSACTION ISOLATION LEVEL SERIALIZABLE;
SELECT PersonID, LastName, FirstName, City, Balance
FROM balances
WHERE City = 'New York'
  AND Balance > 5000
ORDER BY Balance DESC
LIMIT 10;