.PHONY: generate build run

all: generate build run

generate:
	@sh generate.sh auth

build:
	go build ./cmd/auth-service

run:
	./auth-service

run-docker:
	docker-compose build
	docker-compose up --remove-orphans