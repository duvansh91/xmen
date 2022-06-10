PACKAGES := $(shell go list ./...)

test: 
	@go test -count=1 -coverpkg=./... -coverprofile=./coverage/coverage.out -timeout 60s $(PACKAGES)
	@go tool cover -html=./coverage/coverage.out -o ./coverage/coverage.html