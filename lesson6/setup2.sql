DROP TABLE IF EXISTS posts CASCADE;
DROP TABLE IF EXISTS comments CASCADE;

create table posts (
  id serial primary key,
  content text,
  author VARCHAR(255)
);

create TABLE comments (
  id serial PRIMARY KEY,
  content text,
  author VARCHAR(255),
  post_id integer references posts(id)
);