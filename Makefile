#!make
include .env

init-dependency:
	go get -u github.com/antonfisher/nested-logrus-formatter
	go get -u github.com/gin-gonic/gin
	go get -u golang.org/x/crypto
	go get -u gorm.io/gorm
	go get -u gorm.io/driver/postgres
	go get -u github.com/sirupsen/logrus
	go get -u github.com/joho/godotenv


migration_up:
	migrate -path app/database/migrations/ -database "mysql://${DB_USER}:${DB_PASS}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}" -verbose up

migration_down:
	migrate -path app/database/migrations/ -database "mysql://${DB_USER}:${DB_PASS}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}" -verbose down 12

migration_fix:
	migrate -path app/database/migrations/ -database "mysql://${DB_USER}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}" force 20240711030318

# Migrate target
migration:
ifndef TABLE_NAME
	$(error TABLE_NAME is undefined. Use 'make migrate TABLE_NAME=<your_table_name>')
endif
	migrate create -ext sql -dir app/database/migrations/ "create_"$(TABLE_NAME)"_table"