#!/bin/sh

goose -dir migrations \
  postgres "user=postgres password=postgres host=localhost port=5432 dbname=ozonmp sslmode=disable" \
  status
