app.setup:
	docker network create --driver bridge pismo_transactions || true
	docker-compose build

app.run:
	docker-compose up app

app.stop:
	docker-compose stop app

db.up:
	docker-compose up -d db

db.stop:
	docker-compose stop db

migrate:
	docker run --rm --network pismo_transactions --env-file .env pismo-transactions-app go run /app/cmd/.

swagger:
	docker run --rm -v .:/app pismo-transactions-app swag init -d /app/api/

generate:
	docker run --rm -v .:/app pismo-transactions-app go generate ./...

test:
	docker run --rm -v .:/app pismo-transactions-app go test `go list ./... | grep -v mock`

test-coverage:
	docker run --rm -v .:/app pismo-transactions-app go test `go list ./... | grep -v mock` -coverprofile cover.out  && go tool cover -html=cover.out