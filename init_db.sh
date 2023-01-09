#!/bin/sh

set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE DATABASE api;
    CREATE USER app WITH ENCRYPTED PASSWORD 'pleasechangeme';
    GRANT ALL PRIVILEGES ON DATABASE app TO api;
EOSQL
