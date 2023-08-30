-- Select DB
USE go_pays_db;
-- Create table
DROP TABLE IF EXISTS accounts;
CREATE TABLE accounts(
  id INT PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(255), 
  balance DECIMAL(15,2) DEFAULT 0,
  createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
  updatedAt DATETIME ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
INSERT INTO accounts(name, balance, createdAt) VALUES('Rafa', 15000.99, NOW());
INSERT INTO accounts(name, balance, createdAt) VALUES('Rayssa', 500, NOW());
INSERT INTO accounts(name, balance, createdAt) VALUES('Dandan', 0, NOW());
