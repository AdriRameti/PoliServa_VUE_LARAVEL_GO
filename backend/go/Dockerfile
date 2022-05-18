FROM golang:1.17-alpine

WORKDIR /app/go

COPY ./go .

COPY ./db/wait-for-it.sh .

RUN chmod +x wait-for-it.sh

RUN go mod tidy

EXPOSE 8003