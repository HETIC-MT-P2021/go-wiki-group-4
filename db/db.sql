CREATE DATABASE gomvc;

CREATE TABLE article
(
  article_id serial PRIMARY KEY,
  title VARCHAR (128) NOT NULL,
  content TEXT NOT NULL,
  UNIQUE (title)
);

CREATE TABLE comment (
    comment_id serial PRIMARY KEY,
    article_id int references article(article_id),
    content VARCHAR (256) NOT NULL
);
