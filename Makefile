install-mysql-db:
	docker-compose -f /home/cmc/Documents/Go/oneclick/docker/db/mysql.yml up

build:
	go build
# lint:
# 	golangci-lint run
# test:
# 	go test -v ./... -tags skipTest
# build:
# 	go build -v ./...
install-grpc-gateway:
	go get \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
		google.golang.org/protobuf/cmd/protoc-gen-go \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc

buf-lint:
	buf lint
buf-generate:
	buf generate
generate-test:
	protoc -I . \
		-I ${GOPATH}/src \
		-I third-party/ \
		--go_out=.\
		--go-grpc_out=.\
		--grpc-gateway_out=. \
		--grpc-gateway_opt logtostderr=true \
		--validate_out="lang=go:./" \
		--proto_path=./proto/categories/v1/categories.proto ./proto/categories/v1/*.proto \

run-local:
	go run main.go serve

connect-cockroach-db-local:
	cockroach sql --insecure --host=localhost:26257

swag-gen:
	swag init