
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

create table users(
  id  SERIAL PRIMARY KEY,
  username varchar(255),
  email varchar(255),
  enabled boolean DEFAULT False,
  password varchar(255),
  last_login timestamp,
  role varchar(255),
  firstname varchar(255),
  lastname varchar(255)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

drop table users;
