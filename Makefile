gen:
	go generate ./...

fmt:
	go fmt ./...

dep:
	go get -u && go mod tidy && go mod vendor && go mod verify

test:
	go test -mod vendor ./...

clean:
	go clean -mod vendor .
