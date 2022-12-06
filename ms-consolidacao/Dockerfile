FROM golang:1.19

WORKDIR /go/app

RUN apt-get update && apt-get install -y librdkafka-dev

EXPOSE 8080
CMD ["tail", "-f", "/dev/null"]