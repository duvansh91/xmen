PACKAGES := $(shell go list ./...)

test: 
	mkdir -p coverage
	@go test -count=1 -coverpkg=./... -coverprofile=./coverage/coverage.out -timeout 60s $(PACKAGES)
	@go tool cover -html=./coverage/coverage.out -o ./coverage/coverage.html

run:
	@go run cmd/main.go