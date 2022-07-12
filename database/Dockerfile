FROM postgres:10.3

COPY up.sql /docker-entrypoint-initdb.d/1.sql

CMD ["postgres"]

# docker build -t db-pg-go .
# docker run -d -p 54321:5432 db-pg-go