DROP TABLE IF EXISTS users;
CREATE TABLE users (
  id       INT AUTO_INCREMENT NOT NULL,
	name     VARCHAR(128) NOT NULL,
	PRIMARY KEY (id)
);

INSERT INTO users
  (name)
VALUES
  ("Sam"),
  ("Afr");

DROP TABLE IF EXISTS tasks;
CREATE TABLE tasks(
  id         INT AUTO_INCREMENT NOT NULL,
  name      VARCHAR(128) NOT NULL,
  status     VARCHAR(255) NOT NULL,
  comment         VARCHAR(255),
  updated_at datetime DEFAULT NULL,
  created_at datetime DEFAULT NULL,
  user_id INT,
  PRIMARY KEY (id),
  FOREIGN KEY (user_id) REFERENCES users(id)
);

INSERT INTO tasks
  (name, status, comment,user_id)
VALUES
  ('Home Work', 'Incomplete', 'corse 30',1),
  ('Office Work', 'In Progress', 'Golang',1),
  ('Office Work', 'In Progress', 'Framework',2);

INSERT INTO tasks
  (name, status, comment,user_id)
VALUES
("Home Work", "Ka chole", "pathao", 2);


