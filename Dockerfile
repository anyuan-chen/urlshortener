# syntax=docker/dockerfile:1
FROM golang:1.18-alpine
WORKDIR /app
EXPOSE 8080
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./
RUN go build -o /docker-gs-ping
CMD ["/docker-gs-ping"]
