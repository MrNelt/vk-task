build:
	docker-compose build

run: build
	docker-compose up -d

destroy:
	docker-compose down

