FROM postgres
# The above image if of architecture AMD64 with Ubuntu on top
COPY ./init/01_init.sql /docker-entrypoint-initdb.d/
COPY ./init/02_init.sql /docker-entrypoint-initdb.d/
