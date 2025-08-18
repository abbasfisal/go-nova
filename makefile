include: .dev.env



upm: ## Run migrations up
	@migrate -path ./migrations/ -database "mysql://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)" up

down-one: ## Rollback last migration (one step)
	@migrate -path ./migrations/ -database "mysql://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)" down -steps 1

down-all: ## Rollback all migrations
	@migrate -path ./migrations/ -database "mysql://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)" down

version: ## Print current migration version
	@migrate -path ./migrations/ -database "mysql://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)" version



docker-up: ## docker compose up
	 @docker compose -f docker-compose.dev.yml --env-file .dev.env  up



help: ## Show this help
	@echo ""
	@echo "Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[0;32m%-15s\033[0m %s\n", $$1, $$2}'
	@echo ""
