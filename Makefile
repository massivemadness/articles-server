.PHONY: deps
deps:
	go mod tidy -v

.PHONY: build-binary
build-binary:
	go build -o ./build/output/main ./cmd/articles-server

.PHONY: build-image
build-image:
	docker build -f build/Dockerfile -t articles-server --progress=plain .

.PHONY: run
run:
	go run ./cmd/articles-server

.PHONY: test
test:
	go test ./...

.PHONY: clean
clean:
	rm -r ./build/output