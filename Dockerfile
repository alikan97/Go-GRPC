FROM golang:alpine as builder

WORKDIR /go/src/app

ENV GO111MODULE=on

RUN go get github.com/alikan97/Go-GRPC

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./run .

FROM alpine:latest

RUN apk --no-cache add ca-certification
WORKDIR /root/

COPY --from=builder /go/src/app/run .
EXPOSE 8080
CMD [ "./run" ]