.PHONY: dev
.PHONY: dev-rebuild

dev:
	docker-compose -f docker-compose.yml up

dev-rebuild:
	docker-compose -f docker-compose.yml up --build

dev-frontend-rebuild:
	docker-compose -f docker-compose.yml up -d --no-deps --build frontend

dev-backend-rebuild:
	docker-compose -f docker-compose.yml up -d --no-deps --build backend
