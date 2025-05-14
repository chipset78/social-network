.DEFAULT_GOAL := start

# Migrations
.PHONY: migrate-up migrate-down migrate-status

migrate-up:
	docker-compose exec app goose -dir /app/migrations postgres "user=postgres dbname=social_network host=db password=postgres sslmode=disable" up

migrate-down:
	docker-compose exec app goose -dir /app/migrations postgres "user=postgres dbname=social_network host=db password=postgres sslmode=disable" down

migrate-status:
	docker-compose exec app goose -dir /app/migrations postgres "user=postgres dbname=social_network host=db password=postgres sslmode=disable" status

# Service commands
.PHONY: init start stop clean

init:
	docker-compose down -v && docker-compose up -d --build
	$(MAKE) migrate-up

start:
	docker-compose up -d

stop:
	docker-compose down

clean:
	docker-compose down -v
