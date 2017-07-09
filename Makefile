ADDRESS?=:8001

.PHONY: build

build:
	@go build -o webapp/webapp -ldflags "-X main.address=${ADDRESS}" webapp/main.go