-- +goose Up
CREATE TABLE if not exists requests (
  id BIGSERIAL PRIMARY KEY,
  service VARCHAR(255) not null,
  "user" VARCHAR(255) not null,
  text TEXT,
  removed BOOLEAN,
  created TIMESTAMP,
  updated TIMESTAMP
);

CREATE TABLE if not exists requests_events (
  id BIGSERIAL PRIMARY KEY,
  request_id BIGINT,
  type INT,
  status INT,
  payload JSONB,
  updated TIMESTAMP,
  FOREIGN KEY (request_id) REFERENCES requests (id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE requests_events;
DROP TABLE requests;
