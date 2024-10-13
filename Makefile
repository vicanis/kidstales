.PHONY: build
build:
	docker compose build

.PHONY: run
run: build
	docker compose up

.PHONY: pull
pull:
	git pull --rebase

.PHONY: upgrade
upgrade: pull build
	docker compose up -d
