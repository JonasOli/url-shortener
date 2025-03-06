build-dev:
	docker-compose -f docker-compose.dev.yml up -d --build

start-dev:
	docker-compose -f docker-compose.dev.yml up -d

logs:
	docker logs -f url-shortener-app

start-prod:
	docker-compose -f docker-compose.yml up -d

stop-dev:
	docker-compose -f docker-compose.dev.yml down

stop:
	docker-compose down
