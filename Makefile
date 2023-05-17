.PHONY: web
web:
	@echo "Running docker-compose up..."
	docker-compose -f standalone-docker-compose.yaml build web
	@echo "Done!"

.PHONY: standalone-up
standalone-up:
	@echo "Running docker-compose up..."
	docker-compose -f standalone-docker-compose.yaml up -d
	@echo "Done!"

.PHONY: standalone-down
standalone-down:
	@echo "Running docker-compose down..."
	docker-compose -f standalone-docker-compose.yaml down
	@echo "Done!"

.PHONY: replicated-up
replicated-up:
	@echo "Running docker-compose up..."
	docker-compose -f replicated-docker-compose.yaml up -d
	@echo "Done!"

.PHONY: replicated-up
replicated-down:
	@echo "Running docker-compose up..."
	docker-compose -f replicated-docker-compose.yaml down
	@echo "Done!"