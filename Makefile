build-api:
	@go build -o api ./cmd/api

build: build-api
	
run: build-api
	@clear && ./api
	
docker-up: 
	@docker compose up -d

docker-down: 
	@docker compose down

docker-delete-dbdata:
	@sudo rm -rf ./.docker/mongo/mongodb_data/

gen-mock-dbrepo:
	@mockgen -source=./internal/repo/db/interface.go -destination=./internal/mock/dbrepo/mock_dbrepo.go -package=mock_dbrepo

gen-mock-uc:
	@mockgen -source=./internal/usecase/base.go -destination=./internal/mock/usecase/mock_usecase.go -package=mock_usecase