CREATE USER 'admin'@'localhost' IDENTIFIED BY 'admin';
GRANT PROCESS, REPLICATION CLIENT ON *.* TO 'admin'@'localhost';
GRANT SELECT ON performance_schema.* TO 'admin'@'localhost';