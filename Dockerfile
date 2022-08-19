FROM golang:1.18-alpine AS build-env

WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY *.go ./
RUN go build -o /factor
EXPOSE 8080
CMD [ "/docker-gs-ping" ]