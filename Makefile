include .env

get:
	go get

run:
	go run main.go

build:
	go build -o app

build_linux:
	GOOS=linux GOARCH=amd64 go build -o app

db_up:
	migrate -path db/ -database "postgresql://$(DBUSER):$(DBPASS)@$(DBHOST):$(DBPORT)/$(DBNAME)?sslmode=disable" -verbose up

db_down:
	migrate -path db/ -database "postgresql://$(DBUSER):$(DBPASS)@$(DBHOST):$(DBPORT)/$(DBNAME)?sslmode=disable" -verbose down