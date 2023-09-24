SET TRANSACTION ISOLATION LEVEL REPEATABLE READ ;
-- alter database db set default_transaction_isolation TO 'SERIALIZABLE';
SELECT current_setting('transaction_isolation');
