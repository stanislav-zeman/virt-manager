.PHONY: lint
lint:
	mkdir -p frontend/dist
	golangci-lint run

.PHONY: dev
dev:
	wails dev

.PHONY: build
build:
	wails build

.PHONY: test
test:
	go test -cover -coverprofile cover.out .../.

