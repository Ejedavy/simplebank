postgres:
	docker run --name bank_db_container -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres

createdb:
	docker exec -it bank_db_container createdb --username root --owner root simple_bank_db

dropdb:
	docker exec -it bank_db_container dropdb simple_bank_db

createmigration:
	migrate create -ext sql -dir ./db/migrations -seq $(name)

migrateup:
	migrate -path ./db/migrations -database "postgresql://root:secret@localhost:5432/simple_bank_db?sslmode=disable" up

migratedown:
	migrate -path ./db/migrations -database "postgresql://root:secret@localhost:5432/simple_bank_db?sslmode=disable" down

test:
	go test ./... -v -cover
.PHONY: postgres createdb dropdb createmigration migrateup migratedown test