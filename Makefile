.PHONY:
lint:
	golangci-lint run

.PHONY:
dev:
	wails dev

.PHONY:
build:
	wails build

.PHONY:
test:
	go test -cover -coverprofile cover.out .../.

