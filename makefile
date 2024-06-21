include .env

create_container:
	@docker run --name ${DB_DOCKER_CONTAINER} -p 5432:5432 -e POSTGRES_USER=${DB_DOCKER_USER} -e POSTGRES_PASSWORD=${DB_DOCKER_PASSWORD} -d postgres:16-alpine

start_container:
	@docker start ${DB_DOCKER_CONTAINER}

create_db:
	@docker exec -it ${DB_DOCKER_CONTAINER} createdb --username=${DB_DOCKER_USER} --owner=${DB_DOCKER_USER} ${DB_NAME}

drop_db:
	@docker exec -it ${DB_DOCKER_CONTAINER} dropdb ${DB_NAME}

open_db:
	@docker exec -it ${DB_DOCKER_CONTAINER} psql -U ${DB_DOCKER_USER}

migrateup:
	@goose -dir sql/schema postgres "postgresql://${DB_DOCKER_USER}:${DB_DOCKER_PASSWORD}@localhost:5432/${DB_NAME}?sslmode=disable" up

migratedown:
	@goose -dir sql/schema postgres "postgresql://${DB_DOCKER_USER}:${DB_DOCKER_PASSWORD}@localhost:5432/${DB_NAME}?sslmode=disable" down
	
run:
	@templ generate
	@npx tailwindcss -o static/css/styles.css --minify
	@go build -o app.exe main.go && app.exe
	
test:
	@go test -v ./...
	