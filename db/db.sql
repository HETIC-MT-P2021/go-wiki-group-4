CREATE DATABASE IF NOT EXISTS gomvc;

USE gomvc;

CREATE TABLE IF NOT EXISTS article
(
  article_id serial PRIMARY KEY,
  title VARCHAR (128) NOT NULL,
  content TEXT NOT NULL,
  UNIQUE (title)
);

CREATE TABLE IF NOT EXISTS comment (
  comment_id serial PRIMARY KEY,
  article_id int references article(article_id),
  content VARCHAR (256) NOT NULL
);
