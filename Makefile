.PHONY: go_lint postgres_run migrate_check_install migrate_up migrate_down swagger_check_install swagger_spec swagger_ui

go_lint:
	docker run --rm -v ${PWD}:/app -w /app/ golangci/golangci-lint:v1.36-alpine golangci-lint run -v --timeout=5m

postgres_run:
	docker run --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=spoon -e PGDATA=/var/lib/postgresql/data/pgdata -v /custom/mount:/var/lib/postgresql/data -d postgres:9.6.20

migrate_check_install:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

migrate_up: migrate_check_install
	migrate -source file://internal/infrastructure/database/migrations/ -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:${POSTGRES_PORT}/cart_api?sslmode=disable" up

migrate_down: migrate_check_install
	migrate -source file://internal/infrastructure/database/migrations/ -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:${POSTGRES_PORT}/cart_api?sslmode=disable" down

swagger_check_install:
	go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger_spec: swagger_check_install
	swagger generate spec -o ./swagger.yml

swagger_ui: swagger_check_install
	swagger serve -F=swagger swagger.yml