build:
	docker-compose build

run: build
	docker-compose up

destroy:
	docker-compose down

