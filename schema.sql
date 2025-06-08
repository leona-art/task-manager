CREATE TABLE tasks (
  id CHAR(36) NOT NULL PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  `description` TEXT NOT NULL,
  kind ENUM('todo', 'progress','issue') NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  UNIQUE (title, kind)
);

CREATE TABLE todos (
  id CHAR(36) NOT NULL PRIMARY KEY,
  FOREIGN KEY (id) REFERENCES tasks(id) ON DELETE CASCADE ON UPDATE CASCADE,
  status ENUM('pending', 'done') NOT NULL DEFAULT 'pending'
);

CREATE TABLE progress (
  id CHAR(36) NOT NULL PRIMARY KEY,
  FOREIGN KEY (id) REFERENCES tasks(id) ON DELETE CASCADE ON UPDATE CASCADE,
  status ENUM('not started', 'in progress', 'completed') NOT NULL DEFAULT 'not started',
  solution TEXT NULL
);

CREATE TABLE issues (
  id CHAR(36) NOT NULL PRIMARY KEY,
  FOREIGN KEY (id) REFERENCES tasks(id) ON DELETE CASCADE ON UPDATE CASCADE,
  status ENUM('open', 'investigating','resolving','closed') NOT NULL DEFAULT 'open',
  solution TEXT NULL,
  cause TEXT NULL
);
