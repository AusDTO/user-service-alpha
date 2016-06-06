FROM postgres

COPY users.sql /users.sql

CMD sleep 5 && psql -h db -U postgres -a -w -f /users.sql
