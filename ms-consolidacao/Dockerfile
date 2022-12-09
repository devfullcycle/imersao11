FROM golang:1.19

WORKDIR /go/app

RUN apt-get update && apt-get install -y librdkafka-dev

RUN go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.15.2

EXPOSE 8080
CMD ["tail", "-f", "/dev/null"]