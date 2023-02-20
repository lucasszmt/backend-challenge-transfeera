install_deps:
	@echo "Installing go deps"
	@go mod tidy

run_migrations:
	@echo "Running migration"
	@go run cmd/migration/migration.go cmd/migration/queries.go
	@echo "migrations finished"

coverage_tests:
	echo "Running tests"
	go clean -testcache
	go test ./... -covermode=atomic -coverprofile=/tmp/coverage.out -coverpkg=./... -count=1
	@#goverreport -coverprofile=/tmp/coverage.out -sort=block -order=desc -threshold=90 || (echo -e "**********Minimum test coverage was not reached(90%)**********"; exit 1)
	go tool cover -html=/tmp/coverage.out

up-database:
	docker compose up database -d

run:
	docker compose up transfeera-backend

init: up-database run_migrations
	@docker compose up transfeera-backend --build

stop:
	@docker compose stop

down: stop
	@docker compose down