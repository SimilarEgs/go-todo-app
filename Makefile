postgres: #to start postgres container
	docker run --name=todo-db -e POSTGRES_PASSWORD='Eretika12' -p 5432:5432 -d --rm postgres

migrateup:
	migrate -path ./migrations -database 'postgres://postgres:Eretika12@localhost:5432/postgres?sslmode=disable' up

migratedown:
	migrate -path ./migrations -database 'postgres://postgres:Eretika12@localhost:5432/postgres?sslmode=disable' down

postgresD:
	docker exec -it todo-db psql -U postgres

.PHONY: postgres postgresD migratedown migrateup