build:
	@go build -o bin/impelsys

run: build
	@./bin/impelsys

