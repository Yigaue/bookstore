-- connect on CLI: mysql -u -p
-- Enter password:
-- mysql> CREATE DATABASE bookstore;
-- mysql> USE bookstore;
-- mysql> SOURCE /path to/bookstore.sql;
-- mysql> exit (closes connection);

DROP TABLE IF EXISTS book;
CREATE TABLE book (
  id         INT AUTO_INCREMENT NOT NULL,
  title      VARCHAR(128) NOT NULL,
  author     VARCHAR(255) NOT NULL,
  price      DECIMAL(5,2) NOT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO book
  (title, author, price)
VALUES
  ("A Day in the Life of Abed Salama", "Nathan Thrall", 90.18),
	("King: A life", "Jonathan Eig", 56.99),
	("Where we go from here", "Bernie Sanders", 23.99),
	("Buiding a dream server", "Yiga ue", 39.99),
	("Clean Code ", "Robert C Martin", 49.99);