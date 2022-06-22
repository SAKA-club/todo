-- migrate:up
CREATE TABLE IF NOT EXISTS item (
                                    id BIGSERIAL NOT NULL PRIMARY KEY,
                                    title text NOT NULL ,
                                    body text,
                                    priority bool,
                                    schedule_time timestamptz,
                                    complete_time timestamptz,
                                    update_time timestamptz,
                                    delete_time timestamptz,
                                    create_time timestamptz NOT NULL DEFAULT now()
);


-- migrate:down

drop table item;