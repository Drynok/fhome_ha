.PHONY: up down build test protos

up:
	docker-compose up -d

down:
	docker-compose down

build:
	docker-compose up -d --build app

test:
	go test ./... -v