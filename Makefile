# ------------------------------ Golang Command ------------------------------ #
start:
	go run ./cmd/

test:
	go test -v -coverprofile="c.out" ./...

coverage:
	go tool cover -html="c.out"




# ---------------------------------- Docker ---------------------------------- #
local:
	echo "Starting local environment"
	docker-compose -f docker-compose.local.yaml up --build -d

down-local:
	echo "Drop local container"
	docker-compose -f docker-compose.local.yaml down

clean:
	docker system prune -f

# --------------------------------- Migration -------------------------------- #
migrateup:
	migrate -path db/migrations -database "postgresql://postgres:postgres@localhost:5444/store_db?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://postgres:postgres@localhost:5444/store_db?sslmode=disable" -verbose down

.PHONY: start test coverage local down-local clean migrateup migratedown