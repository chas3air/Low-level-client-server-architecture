build-up:
	@docker compose up --build -d

down:
	@docker compose down

refresh:
	@docker compose down
	@docker compose up --build -d

COMMIT_NAME ?= refactoring

# make COMMIT_NAME="..." fast_commit
fast_commit:
	@git add .
	@git commit -m "$(COMMIT_NAME)"