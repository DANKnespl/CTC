# syntax=docker/dockerfile:1

FROM golang:1.22

WORKDIR /app

COPY go.sum ./

COPY go.mod ./
RUN go mod download

COPY *.yaml ./
COPY *.go ./

RUN go build -o /gas-pump

CMD ["/gas-pump" ]