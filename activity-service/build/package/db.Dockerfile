FROM postgres:17
RUN echo 'CREATE EXTENSION IF NOT EXISTS "uuid-ossp";' > /docker-entrypoint-initdb.d/init.sql
