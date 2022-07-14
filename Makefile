install-mysql-db:
	docker-compose -f /home/cmc/Documents/Go/oneclick/docker/db/mysql.yml up
lint:
	golangci-lint run
test:
	go test -v ./... -tags skipTest
build:
	go build -v ./...
grpc-client:
	grpc-client-cli --protoimports --proto ./proto/metronion/v1/metronion.proto localhost:9000
install-evans-cli:
	brew tap ktr0731/evans
	brew install evans
evans-cli:
	evans -r repl -p 9000
evans-cli-dev:
	evans --host dev-grpc.metrogalaxy.io --port 443 -r repl
grpcurl-dev:
	grpcurl dev-grpc.metrogalaxy.io:443 list
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
# generate:
# 	protoc -I . \
# 		-I ${GOPATH}/src \
# 		-I third-party/ \
# 		--proto_path=./api/v1/protobuf metronion.proto \
# 		--go_out=.\
# 		--go-grpc_out=.\
# 		--grpc-gateway_out=. \
# 		--grpc-gateway_opt logtostderr=true \
# 		--validate_out="lang=go:./" \
# 		--proto_path=./api/v1/protobuf ./api/v1/protobuf/*.proto \

run-local:
	godotenv -f local.env go run main.go serve

publish-wearable-metadata-local:
	godotenv -f local.env go run main.go publish-wearable-metadata --file metronion_wearables.csv

publish-metronion-metadata-local:
	godotenv -f local.env go run main.go publish-metadata --file data.csv

publish-wearable-metadata-dev:
	godotenv -f dev.env go run main.go publish-wearable-metadata --file metronion_wearables.csv

publish-metronion-metadata-dev:
	godotenv -f dev.env go run main.go publish-metadata --file data.csv

start-cockroach-db-local:
	cockroach start-single-node --advertise-addr 'localhost' --http-addr=localhost:8081 --insecure

connect-cockroach-db-local:
	cockroach sql --insecure --host=localhost:26257

init-cockroach-db-local:
	cockroach sql --insecure --host=localhost:26257 -f ./sql/init_database.sql
init-cockroach-db-dev:
	cockroach sql --url "postgresql://metroadmin@free-tier8.aws-ap-southeast-1.cockroachlabs.cloud:26257/defaultdb?sslmode=verify-full&options=--cluster=dev-metrogalaxy-1716" -f ./sql/init_database.sql

start-grpcweb-proxy:
	grpcwebproxy \
		--server_tls_cert_file=./docker/grpcwebproxy/server-cert.pem \
		--server_tls_key_file=./docker/grpcwebproxy/server-key.pem \
		--backend_addr=localhost:9000 \
		--server_http_debug_port 8081 \
		--backend_tls_noverify \
		--server_http_max_read_timeout=30s \
		--server_http_max_write_timeout=30s \
		--allow_all_origins \
		--enable_request_debug