default:
	build
clean:
	rm -rf build
	rm diskusage

setup:
	go get -u "github.com/golang/lint/golint"

test:
	go test

lint:
	golint .

build-win-64:
	GOOS=windows GOARCH=amd64 go build -o build/windows/64bit/diskusage.exe

build-win-32:
	GOOS=windows GOARCH=386 go build -o build/windows/32bit/diskusage.exe

build-lin-64:
	GOOS=linux GOARCH=amd64 go build -o build/linux/64bit/diskusage

build-lin-32:
	GOOS=linux GOARCH=386 go build -o build/linux/32bit/diskusage

build-mac-64:
	GOOS=darwin GOARCH=amd64 go build -o build/mac/64bit/diskusage

build-mac-32:
	GOOS=darwin GOARCH=386 go build -o build/mac/64bit/diskusage

build: build-win-64 build-win-32 build-lin-64 build-lin-32 build-mac-64 build-mac-32

help:
	@echo "COMMANDS:"
	@echo "  clean          Remove all generated files."
	@echo "  setup          Setup development environment."
	@echo "  test           Run tests."
	@echo "  lint           Run analysis tools."
	@echo "  build          Build all architectures for all operating systems"