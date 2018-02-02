-- +goose Up
-- SQL in this section is executed when the migration is applied.

insert into users (username, email, enabled, last_login, role) values ('admin', 'arijeet.baruah@weboniselab.com', true, CURRENT_TIMESTAMP, 'admin');

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
