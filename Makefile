proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./pb/auth/auth.proto
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./pb/events/events.proto

build:
	go build -o ./api/main.exe ./api/cmd/main.go
	go build -o ./events/main.exe ./events/cmd/main.go

run-api-gateway:
	./api/main.exe

run-events-service:
	./events/main.exe

dev-api-gateway:
	cd ./api && air

dev-events-service:
	cd ./events && air