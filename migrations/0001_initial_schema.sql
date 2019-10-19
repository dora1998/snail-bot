-- +migrate Up
CREATE TABLE tasks (
                      id varchar(36) NOT NULL PRIMARY KEY,
                      body text NOT NULL,
                      deadline datetime NOT NULL,
                      created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                      created_by text
);

-- +migrate Down
DROP TABLE tasks;