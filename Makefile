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
migrate_up:
	migrate -path db/migrations -database "postgresql://postgres:postgres@localhost:5444/store_db?sslmode=disable" -verbose up

migrate_down:
	migrate -path db/migrations -database "postgresql://postgres:postgres@localhost:5444/store_db?sslmode=disable" -verbose down

# ----------------------------- SSL/TLS commands ----------------------------- #
#generate private key self-signed certificate (public key)
gen_private_key:
	openssl genrsa -out server.key 2048
	openssl ecparam -genkey -name secp384r1 -out server.key

#generate self-signed certificate (public key)
gen_self_signed_cert:
	openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650

.PHONY: start test coverage local down-local clean migrate_up migrate_down gen_private_key gen_self_signed_cert

