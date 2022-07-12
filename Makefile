postgres: #to start postgres container
	docker run --name=todo-db -e POSTGRES_PASSWORD='postgres' -p 5432:5432 -d postgres

migrateup:
	migrate -path ./migrations -database 'postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable' up

migratedown:
	migrate -path ./migrations -database 'postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable' down

postgresD:
	docker exec -it todo-db psql -U postgres

.PHONY: postgres postgresD migratedown migrateup
