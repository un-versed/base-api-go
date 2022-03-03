######## Stage 1 - builder #######
FROM golang:latest as builder

WORKDIR /app
COPY . .

ENV GO111MODULE=on
ENV SERVER_PORT=8080

RUN make build-for-docker

######## Stage 2 - runner #######
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]