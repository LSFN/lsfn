generate:
	protoc -I . api/proto/*.proto --go_out=plugins=grpc:.
