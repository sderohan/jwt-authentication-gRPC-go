gen: 
	protoc --proto_path=proto proto/*.proto  --go_out=plugins=grpc:pb

client:
	go run cmd/client/main.go

server:
	go run cmd/server/main.go
