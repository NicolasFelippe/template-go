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

.PHONY: sqlc migrateup migrateup1 migratedown migratedown1 create-migration patch minor major