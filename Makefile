build:
	go build -o ./build/output/main ./cmd/articles-server

run:
	go run ./cmd/articles-server

clean:
	rm -r ./build/output