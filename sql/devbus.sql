CREATE TABLE IF NOT EXISTS devbus_user (
  id        SERIAL4 PRIMARY KEY,
  name      VARCHAR(36),
  email     VARCHAR(128),
  salt      VARCHAR(20),
  password  VARCHAR(64),
  full_name VARCHAR(64)
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
