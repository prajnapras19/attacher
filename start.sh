#!/bin/sh

cat << EOF > .env
MYSQL_MIGRATIONS_FOLDER=/app/migrations
EOF

make mysql-migrate-up
if [ $? -ne 0 ]; then
  echo "Make command failed"
  exit 1
fi

./main