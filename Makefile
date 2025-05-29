CURRENT_DIR=$(shell pwd)
DATABASE_URL="postgres://quoter:quoter@localhost:5432/quoter_db?sslmode=disable" # this is for local development purpose only!

run:
	go run cmd/main.go

run-local-docker-up:
	docker-compose -f docker-compose.local.yaml up --build -d

new-migration:
	migrate create -ext sql -dir ${CURRENT_DIR}/migrations -seq -digits 8 $(name)

migrate:
	migrate -source file://${CURRENT_DIR}/migrations -database ${DATABASE_URL} -verbose up

.PHONY: migrate