#!/bin/bash

# Wait for MySQL server to come up
until mysql -u${BOT_DB_USERNAME} -p${BOT_DB_PASSWORD} -h${BOT_DB_HOST} -P${BOT_DB_PORT} -e "SELECT 1"; do sleep 1; done

# Start server
exec "$@"