build:
	go build -o ./bin/pin ./cmd/pin

run: 
	./bin/pin

all: build run  
