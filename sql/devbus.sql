CREATE TABLE IF NOT EXISTS devbus_user (
  id          SERIAL4 PRIMARY KEY,
  name        VARCHAR(36)  NOT NULL UNIQUE,
  email       VARCHAR(128) NOT NULL UNIQUE,
  salt        VARCHAR(20)  NOT NULL,
  password    VARCHAR(64)  NOT NULL,
  full_name   VARCHAR(64)  NOT NULL,
  is_activate BOOLEAN      NOT NULL
);

CREATE TABLE IF NOT EXISTS project (
  id          SERIAL4 PRIMARY KEY,
  name        VARCHAR(256),
  description TEXT,
  desc_type   VARCHAR(8)
);

CREATE TABLE IF NOT EXISTS version (
  id          SERIAL4 PRIMARY KEY,
  name        VARCHAR(128),
  description TEXT,
  desc_type   VARCHAR(8)
);

CREATE TABLE IF NOT EXISTS issue (
  id          SERIAL8 PRIMARY KEY,
  state       SMALLINT,
  title       VARCHAR(256),
  body        TEXT,
  body_type   VARCHAR(8),
  create_time TIMESTAMP,
  update_time TIMESTAMP
);

CREATE TABLE IF NOT EXISTS comment (
  id           SERIAL8 PRIMARY KEY,
  content      TEXT,
  content_type VARCHAR(8)
);
