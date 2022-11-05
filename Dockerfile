FROM golang:1.19.2-alpine3.16 AS builder

RUN apk update \
    && apk add openssl
    
WORKDIR /app
COPY . .
RUN go build -o ./out .

FROM alpine:latest
WORKDIR /app

RUN apk --no-cache add ca-certificates
COPY --from=builder /app/out .
COPY --from=builder /app/keyfiles .
COPY --from=builder .env .

RUN echo $(ls .)

EXPOSE 8081
CMD ["./out"]