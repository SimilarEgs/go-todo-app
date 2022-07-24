DB_URL=postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable

postgres:
	docker run --name todo-db --network todo-app-network -e POSTGRES_PASSWORD='postgres' -p 5432:5432 -d postgres:14-alpine

network:
	docker network create todo-app-network
	
migrate-up:
	migrate -path ./migrations -database "$(DB_URL)" -verbose up

migrate-down:
	migrate -path ./migrations -database "$(DB_URL)" -verbose down

compose-up:
	docker compose up --build

compose-down:
	docker compose down

swag:
	swag init -g pkg/handler/router.go

swag-run:
	DISABLE_SWAGGER_HTTP_HANDLER='' GIN_MODE=debug CGO_ENABLED=0 go run -tags migrate ./cmd/

.PHONY:  network postgres  migratedown migrateup compose-up compose-down swag swag-run