postgres:
	docker run --name todo-db --network app-network -e POSTGRES_PASSWORD='postgres' -p 5432:5432 -d postgres:14-alpine

network:
	docker network create todo-app-network
	
opendb:
	docker exec -it todo-db psql -U postgres

migrateup:
	migrate -path ./migrations -database 'postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable' up

migratedown:
	migrate -path ./migrations -database 'postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable' down

todo-app:
	docker run --name test-app --network app-network -p 8080:8080 web-app


.PHONY:  network postgres postgresD migratedown migrateup