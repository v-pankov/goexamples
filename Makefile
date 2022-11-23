all: prepare-bin clean-chat build-chat

clean-chat:
	rm -f bin/chat

build-chat:
	go build -o bin/chat ./cmd/chat

prepare-bin:
	mkdir -p bin

run-chat:
	go run ./cmd/chat/main.go