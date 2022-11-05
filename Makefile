PORT = 5432
PWD = $(shell pwd)
ACCPATH = $(PWD)/

gen-proto:
	protoc --go_out=. --go_opt=Mproto/crypto.proto=example.com/grpc/crypto --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/crypto.proto

clean-proto:
	rm proto/*.go

build:
	go build -o ./out .

run:
	go run .

new-migration:
	migrate create -ext sql -dir migrations/ -seq -digits 5 $(filename)

# migrate:
# 	migrate -source file://migrations/ -database "postgres://postgres:abcd1234@localhost:$(PORT)/Crypto?sslmode=disable" -verbose up
migrate-up:
	migrate -source file://migrations/ -database "postgres://postgres:$(PWD)@crypto.c2f0vwowpfvy.us-east-1.rds.amazonaws.com:$(PORT)/postgres?sslmode=disable" up

swagger:
	GO111MODULE=off swagger generate spec -o ./swagger.yaml --scan-models

## export PATH=$PATH:/usr/local/go/bin

create-keypair:
	@echo "Creating an rsa 256 key pair"
	openssl genpkey -algorithm RSA -out ./rsa_private.pem -pkeyopt rsa_keygen_bits:2048
	openssl rsa -in ./rsa_private.pem -pubout -out ./rsa_public.pem