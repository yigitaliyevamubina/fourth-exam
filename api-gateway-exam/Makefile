CURRENT_DIR=$(shell pwd)
DB_URL="postgres://postgres:mubina2007@localhost:5432/socialdb?sslmode=disable"

migrate-up:
	migrate -path migrations -database "$(DB_URL)" -verbose up

migrate-down:
	migrate -path migrations -database "$(DB_URL)" -verbose down 1

migrate-force:
	migrate -path migrations -database "$(DB_URL)" -verbose force 3

migrate-file:
	migrate create -ext sql -dir migrations/ -seq insertion_to_tables

proto-gen:
	./scripts/gen-proto.sh ${CURRENT_DIR}

swag-gen:
	~/go/bin/swag init -g ./api/router.go -o api/docs

g:
	go run cmd/main.go