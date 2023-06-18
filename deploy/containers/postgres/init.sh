psql -U "${POSTGRES_USER}" -c "CREATE DATABASE \"users\";"
psql -U "${POSTGRES_USER}" -d users -c "CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";";
