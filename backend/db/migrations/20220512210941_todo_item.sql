-- migrate:up
create table todo_item (
                       id bigserial not null,
                       name varchar(255),
                       email varchar(255) not null
);

-- migrate:down

