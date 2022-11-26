start:
	@go run main.go

build:
	@go build main.go

run:
	@./main

win:
	@go build main.go
	@./main.exe