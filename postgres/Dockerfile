FROM postgres:14-alpine

WORKDIR /app

COPY ./tables/init.sh ./docker-entrypoint-initdb.d/
COPY ./tables/users.sql ./docker-entrypoint-initdb.d/users.sql

CMD ["docker-entrypoint.sh", "postgres"]