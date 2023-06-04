run:
	go run cmd/app/main.go
dev:
	nodemon --exec go run cmd/app/main.go --signal SIGTERM
test:
	go test -v ./...