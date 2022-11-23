all: clean build-chat

clean:
	rm -f browurls

build-chat:
	go build -o chat ./cmd/chat

run-chat:
	go run ./cmd/chat/main.go