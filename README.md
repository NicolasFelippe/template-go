# Template Go
Architecture hexagonal

go v1.19

Libs:
1. Sqlc - Db         https://github.com/kyleconroy/sqlc
2. Gin - Http        https://github.com/gin-gonic/gin
3. Gqlgen - Graphql  https://github.com/99designs/gqlgen
4. Logger ZAP -      https://github.com/uber-go/zap
5. Config Viper -    https://github.com/spf13/viper
6. Bcrypt
7. JWT -             https://github.com/golang-jwt/jwt
8. Paseto -          https://github.com/o1egl/paseto
9. Uuid  -           https://github.com/google/uuid



### Run devlopment on Linux (WSL2)

1. Clone this repository
2. Install go v1.19+

Execute in terminal:
```shell
go mod tidy
make server
```

### Docs api Insominia
 ./docs/api

### Graphql playground
http://localhost:8080/graph/v1

### API Http playground
http://localhost:8080/api/v1


### Run docker

1. docker network create template-network
2. docker run --name postgres12 --network template-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=ORiBLEcTUrdS -d postgres:12-alpine
3. docker build -t templatego:latest .
4. docker run --name template_go --network template-network -p 8080:8080 -e DB_SOURCE="postgresql://root:ORiBLEcTUrdS@postgres12:5432/template_go?sslmode=disable" -e  GIN_MODE=release  templatego:latest


### Migrations
    Run the make command to generate the migration file in 
    the db/migrations folder, after adding table creation 
    commands to the 'UP' and 'DOWN' file, add a command to remove.

1. Install Migrate https://github.com/golang-migrate/migrate/tree/master/cmd/migrate

```shell
sudo curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add -
sudo echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list
sudo apt-get update
sudo apt-get install -y migrate
```

```shell
make create-migration arg=<name>
```
#### Commands migrations

Create file
```shell
make create-migration arg=<name>
```

Submit migrations - Migrations execute on play API
```shell
make migrateup
```

Revert migrations
```shell
make migratedown
```

Submit one migration - Migrations execute on play API
```shell
make migrateup1
```

Revert one migration
```shell
make migratedown1
```


### Util
Generate random 32 bytes
```shell
openssl rand -hex 64 | head -c 32
```