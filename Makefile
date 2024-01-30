postgres:
	docker run --name postgres16 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16-alpine

docker_start:
	docker start postgres16

docker_stop:
	docker stop postgres16

createdb:
	docker exec -it postgres16 createdb --username=root --owner=root bankProject

dropdb:
	docker exec -it postgres16 dropdb bankProject

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bankProject?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bankProject?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown test docker_start docker_stop