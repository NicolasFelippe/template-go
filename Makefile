DB_URL						=	postgresql://root:ORiBLEcTUrdS@localhost:5432/go_template?sslmode=disable
home				        = 	$(shell home)
software_version	  		=	$(shell cat VERSION)
version_array		   		=	$(subst ., ,$(software_version))
major				        = 	$(word 1,${version_array})
minor				        = 	$(word 2,${version_array})
patch				        = 	$(word 3,${version_array})
pwd 				        = 	$(shell pwd)

patch:
	- @echo "BUMPING PATCH"
	- @echo "Current Version: $(software_version)"
	- $(eval patch=$(shell echo $$(($(patch)+1))))
	- @echo "New Version: $(major).$(minor).$(patch)"
	- @printf $(major).$(minor).$(patch) > VERSION

minor:
	- @echo "BUMPING MINOR"
	- @echo "Current Version: $(software_version)"
	- $(eval minor=$(shell echo $$(($(minor)+1))))
	- @echo "New Version: $(major).$(minor).0"
	- @printf $(major).$(minor).0 > VERSION

major:
	- @echo "BUMPING MAJOR"
	- @echo "Current Version: $(software_version)"
	- $(eval major=$(shell echo $$(($(major)+1))))
	- @echo "New Version: $(major).0.0"
	- @printf $(major).0.0 > VERSION

postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=ORiBLEcTUrdS -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root go_template

mock:
	mockgen -package mockdb  -destination=./mocks/sqlc/mock.store.go template-go/internal/sqlc/repositories Store
	mockgen -package userrepositorymock  -destination=./mocks/repositories/mock.userrepository.go template-go/internal/core/domain/users UserRepository
	mockgen -package cryptomock  -destination=./mocks/pkg/crypto/mock.crypto.go template-go/pkg/crypto Crypto
	mockgen -package uidgenmock  -destination=./mocks/pkg/uidgen/mock.uidgen.go template-go/pkg/uidgen UIDGen

gqlgen:
	go run github.com/99designs/gqlgen generate

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

dockerclear:
	docker compose down && docker rmi template-go_api

docker-build:
	docker build -t template_go:latest .

docker-run:
	docker run --name template_go --network template-network -p 8080:8080 -e DB_SOURCE="$(DB_URL)" -e GIN_MODE=release templatego:latest

server:
	go run cmd/http/main.go

test:
	go test -v -cover ./...

.PHONY: sqlc mock server dockerclear gqlgen mock1 migrateup migrateup1 migratedown migratedown1 create-migration patch minor major test