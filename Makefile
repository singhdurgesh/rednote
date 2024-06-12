# note: call scripts from /scripts
include .env
export

run:
	go run main.go server

start_worker:
	go run main.go worker

docker_rm:
	docker-compose rm -f

docker_down:
	docker-compose down

docker_build:
	docker-compose up --build -d

docker_app:
	docker run -it --rm --network=rednote_rednote_network -e "REDNOTE_ENV=docker" -p 8080:8080 rednote-rednote

docker_up:
	docker-compose up -d

docker_build_up: docker_down docker_rm docker_build

# Database Migrations with Goose
# https://github.com/pressly/goose
migration_status:
	goose -dir db/migrations status

migration_create:
	goose -dir db/migrations create $(name) sql

migrate: update_schema
	goose -dir db/migrations up

rollback: update_schema
	goose -dir db/migrations down

# Re-run the latest migration
migration_redo:
	goose -dir db/migrations redo

# Roll back all migrations
migration_reset:
	goose -dir db/migrations reset

migration_version:
	goose -dir db/migrations version

# Check migration files without running them
migration_validate:
	goose -dir db/migrations validate

# -s or --schema-only means dump only the schema, no data
update_schema:
	PGPASSWORD=$(PG_PASSWORD) pg_dump -s -U $(PG_USERNAME) -h $(PG_HOST) -d $(DB_NAME) > db/schema.sql
