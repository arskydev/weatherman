FROM postgres:latest
ENV POSTGRES_DB weatherman
COPY init.sql /docker-entrypoint-initdb.d/
