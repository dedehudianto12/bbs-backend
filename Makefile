include .env
export

DB_URL = postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)
APP=cmd/api/main.go

run:
	air

dev:
	go run cmd/api/main.go

build:
	go build -o bin/api cmd/api/main.go

create-migration:
	migrate create -ext sql -dir migrations -seq $(name)

migrate-up:
	migrate -path migrations -database "$(DB_URL)" up

migrate-down:
	migrate -path migrations -database "$(DB_URL)" down

migrate-version:
	migrate -path migrations -database "$(DB_URL)" version

test:
	go test ./...