build:
	go build -o bin/max-in-nodes

run: build
	./bin/max-in-nodes -numofnodes 100