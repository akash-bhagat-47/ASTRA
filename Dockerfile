FROM postgres:latest AS postgres-setup

ENV POSTGRES_USER=akash.b
ENV POSTGRES_PASSWORD=""
ENV POSTGRES_DB=weatherStation

COPY init.sql /docker-entrypoint-initdb.d/

FROM alpine:latest

WORKDIR /app

COPY ./weatherAppLinux /app/main

RUN apk add --no-cache postgresql-client

EXPOSE 8080

CMD ["./main"]
