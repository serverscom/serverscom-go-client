test:
	go test ./...

deps:
	go mod tidy
	go mod vendor

gen:
	ruby gen.rb
	gofmt -w pkg
