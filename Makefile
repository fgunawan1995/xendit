coverage.out: $(shell find . -type f -print | grep -v vendor | grep "\.go")
	@go test -cover -coverprofile ./coverage.out.tmp ./...
	@cat ./coverage.out.tmp | grep -v '.pb.go' | grep -v 'mock_' > ./coverage.out
	@rm ./coverage.out.tmp

test: coverage.out

cover: coverage.out
	@echo ""
	@go tool cover -func ./coverage.out

cover-html: coverage.out
	@go tool cover -html=./coverage.out

clean:
	@rm ./coverage.out

build:
	@go build

run:
	@go run main.go