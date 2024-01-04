run: build
	@./bin/gopokemon

build:
	@go build -o bin/gopokemon cmd/main/main.go

test:
	@ go test ./...

seed:
	@go run cmd/seed/main.go

drop:
	@go run cmd/drop/main.go

select:
	@go run cmd/select/main.go $(name)
