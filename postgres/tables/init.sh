#!/bin/sh

psql -U $POSTGRES_USER -d $POSTGRES_DB -a -f /app/docker-entrypoint-initdb.d/users.sql