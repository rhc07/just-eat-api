run:
	@go build -o just-eat-api .
	./just-eat-api

test:
	@go test ./...