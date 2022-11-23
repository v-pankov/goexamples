all: clean build-chat

clean:
	rm -f chat

build-chat:
	go build -o chat ./cmd/chat

run-chat:
	go run ./cmd/chat/main.go