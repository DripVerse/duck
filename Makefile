help:
	@echo "Duck Makefile (Requires Go v1.16+)"
	@echo "'help' - Displays the command usage"
	@echo "'build' - Builds the application into a binary/executable" 
	@echo "'install' - Installs the application"
	@echo "'build-windows' - Builds the application for Windows platforms"
	@echo "'build-darwin' - Builds the application for MacOSX platforms"
	@echo "'build-linux' - Builds the application for Linux platforms"
	@echo "'build-all' - Builds the application for all platforms"

build:
	@echo Compiling Duck
	@go build .
	@echo Compile Complete. Run './duck(.exe)'

install:
	@echo Installing Duck
	go install .
	@echo install Complete. Run 'duck'.

build-windows:
	@echo Cross Compiling Duck for Windows x86
	@GOOS=windows GOARCH=386 go build -o ./bin/duck-windows-x32.exe
	@echo Cross Compiling Duck for Windows x64
	@GOOS=windows GOARCH=amd64 go build -o ./bin/duck-windows-x64.exe

build-darwin:
	@echo Cross Compiling Duck for MacOSX x64
	@GOOS=darwin GOARCH=amd64 go build -o ./bin/duck-darwin-x64

build-linux:
	@echo Cross Compiling Duck for Linux x32
	@GOOS=linux GOARCH=386 go build -o ./bin/duck-linux-x32
	@echo Cross Compiling Duck for Linux x64
	@GOOS=linux GOARCH=amd64 go build -o ./bin/duck-linux-x64
	@echo Cross Compiling Duck for Linux Arm32
	@GOOS=linux GOARCH=arm go build -o ./bin/duck-linux-arm32
	@echo Cross Compiling Duck for Linux Arm64
	@GOOS=linux GOARCH=arm64 go build -o ./bin/duck-linux-arm64

build-all: build-windows build-darwin build-linux
	@echo Cross Compiled Duck for all platforms

