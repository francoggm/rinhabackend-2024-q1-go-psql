up: down
	@echo Starting services
	docker-compose up --build -d

up-debug: down
	@echo Starting debug services
	docker-compose -f docker-compose.debug.yml up -d && sleep 1
	cd src && go run .

down:
	@echo Stoping services
	docker-compose down

upd:
	cd src && go mod tidy