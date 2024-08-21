all: run

build:
	@go build -o ./bin/

run: build
	@./bin/expense_tracker