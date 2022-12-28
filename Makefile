DB_URL=postgresql://root:ORiBLEcTUrdS@localhost:5432/go_template?sslmode=disable

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root go_template

sqlc:
	sqlc -f ./configs/sqlc.yaml generate

create-migration:
	migrate create -ext sql -dir db/migration -seq $(arg)

migrateup:
	migrate -path db/migration -database "$(DB_URL)" --verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" --verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" --verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" --verbose down 1

.PHONY: sqlc migrateup migrateup1 migratedown migratedown1 create-migration