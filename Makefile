install-mysql-db:
	docker-compose -f /home/cmc/Documents/Go/oneclick/docker/db/mysql.yml up
lint:
	golangci-lint run
# test:
# 	go test -v ./... -tags skipTest
# build:
# 	go build -v ./...
# install-grpc-gateway:
# 	go get \
# 		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
# 		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
# 		google.golang.org/protobuf/cmd/protoc-gen-go \
# 		google.golang.org/grpc/cmd/protoc-gen-go-grpc

# buf-lint:
# 	buf lint
# buf-generate:
# 	buf generate

run-local:
	go run main.go serve

connect-cockroach-db-local:
	cockroach sql --insecure --host=localhost:26257
swag-gen:
	swag init -g oneclick/api/http/controller/categories_controller.go --output oneclick/api/http/controller/