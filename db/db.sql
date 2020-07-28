CREATE DATABASE IF NOT EXISTS gomvc;

USE gomvc;

CREATE TABLE IF NOT EXISTS article
(
  article_id bigint(20) AUTO_INCREMENT,
  title VARCHAR (128) NOT NULL,
  content TEXT NOT NULL,

  PRIMARY KEY (article_id)
);

CREATE TABLE IF NOT EXISTS comment (
  comment_id bigint(20) AUTO_INCREMENT PRIMARY KEY,
  article_id bigint(20) NOT NULL,
  content VARCHAR (256) NOT NULL,

  FOREIGN KEY (article_id) REFERENCES article(article_id)
);
