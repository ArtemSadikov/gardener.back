# Use the official PostgreSQL base image
FROM postgres:14.2

# Copy the shell script to the container
COPY init.sh /docker-entrypoint-initdb.d/

# Grant execute permissions to the script
RUN chmod +x /docker-entrypoint-initdb.d/init.sh
