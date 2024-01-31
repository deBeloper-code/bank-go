docker-up:
	docker-compose up -d
docker-down:
	docker-compose down
docker-start:
	docker-compose start
docker-restart:
	docker-compose restart
docker-stop:
	docker-compose stop
db-shell:
	docker-compose exec db psql -U root bank_db
createdb:
	docker exec -it bankGo createdb --username=root --owner=root bank_db
dropdb:
	docker exec -it bankGo dropdb bank_db
migratecreate:
	migrate create -ext sql -dir internal/adapter/database/migrations -seq $(name)
migrateup:
	migrate -path internal/adapter/database/migrations -database "postgres://root:secret@localhost:5432/bank_db?sslmode=disable" -verbose up
migratedown:
	migrate -path internal/adapter/database/migrations -database "postgres://root:secret@localhost:5432/bank_db?sslmode=disable" -verbose down
sqlc: 
	sqlc generate
test: 
	go test -v -cover ./...
serve: 
	go run cmd/app/main.go

.PHONY: postgres createdb dropdb migrateup sqlc serve test docker-up docker-down docker-start docker-restart docker-stop db-shell containerDB migratedown migratecreate