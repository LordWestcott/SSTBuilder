run: build
	@./bin/sstbuilder.exe

build:
	@go build -o bin/sstbuilder.exe