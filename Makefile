proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./pb/auth/auth.proto
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./pb/events/events.proto

build:
	go build -o ./api/main.exe ./api/cmd/main.go
	go build -o ./auth/main.exe ./auth/cmd/main.go
	go build -o ./events/main.exe ./events/cmd/main.go

run-api:
	./api/main.exe

run-events:
	./events/main.exe

run-auth:
	./auth/main.exe

dev-api:
	cd ./api && air

dev-events:
	cd ./events && air

dev-auth:
	cd ./auth && air