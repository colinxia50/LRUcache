gen:
	protoc --proto_path=proto --go_out=. --go-grpc_out=. proto/*.proto
clean:
	rm pb/*.go

setk:
	curl -H "Content-Type: application/json" -X POST -d '{"key": "user", "userid":100,"description":"json结构没要求有key就行" }' "127.0.0.1:9000"

getk1:
	curl -i "127.0.0.1:9001?key=user"

getk2:
	curl -i "127.0.0.1:9002?key=user"