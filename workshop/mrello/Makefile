include .env

migrate-up:
	@echo "Migrating up..."
	migrate -path ./db/migrations -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" up

migrate-down:
	@echo "Migrating down..."
	migrate -path ./db/migrations -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" down 1

migrate-create:
	@echo "Creating migration..."
	migrate create -ext sql -dir ./db/migrations -seq $(NAME)

run:
	@echo "Running server..."
	godotenv -f .env go run ./cmd/
