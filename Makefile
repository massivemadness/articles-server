.PHONY: deps
deps:
	go mod tidy -v

.PHONY: build
build:
	go build -o ./build/output/main ./cmd/articles-server

.PHONY: build-image
build-image:
	docker build -t articles-server --progress=plain .

.PHONY: run
run:
	go run ./cmd/articles-server

.PHONY: clean
clean:
	rm -r ./build/output