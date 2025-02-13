default:
	make down
	make gen
	make up

up:
	cd infra && docker compose up --build -d

down:
	cd infra && docker compose --profile loadtest down

init:
	cometbft init --home ./cmt-home

testnet:
	cometbft testnet --n=3 --v=4 --config ./infra/config_template.toml --o=./testnet-home --starting-ip-address 192.167.10.2

gen:
	cd db && sqlc generate
	protoc --go_out=./proto_gen --go-grpc_out=./proto_gen ./audius.proto
	go mod tidy

deps:
	brew install protobuf
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/cometbft/cometbft/cmd/cometbft@v0.38.9

lt:
	@if [ "$(shell docker ps -q -f name=moshpit)" ]; then \
			docker stop moshpit; \
	fi
	@if [ "$(shell docker ps -a -q -f name=moshpit)" ]; then \
			docker rm moshpit; \
	fi
	cd infra && docker compose --profile loadtest up --build -d --no-recreate
	open http://localhost:8080/

mosh:
	docker start moshpit
	open http://localhost:8080/

chill:
	docker stop moshpit

test:
	go test -count=1 ./...
