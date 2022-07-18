#!/bin/sh
# wait-for-postgres.sh

set -e

DB_SOURCE="postgres://postgres:postgres@postgres:5432/postgres?sslmode=disable"

echo "run db migration"
echo 
/app/migrate -path /app/migrations -database "$DB_SOURCE" -verbose up

echo "start the app"
exec "$@"