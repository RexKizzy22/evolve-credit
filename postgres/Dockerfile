FROM postgres:14-alpine

WORKDIR /app

COPY ./tables/init.sh /docker-entrypoint-initdb.d
COPY ./tables/users.sql ./tables/users.sql
