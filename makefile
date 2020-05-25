VERSION=0.0.1

all: test build

.PHONY: test
test:
	go test -v -coverprofile=coverage.out ./...

.PHONY: build
build: build_linux build_darwin build_windows

.PHONY: build_linux
build_linux: nix386 nixamd64

.PHONY: nix386
nix386:
	@GOOS=linux GOARCH=386 go build -o bin/modversion modversion.go
	@cd bin;tar -czf modversion_v${VERSION}_linux_i386.tar.gz modversion
	@rm -f modversion
	
.PHONY: nixamd64
nixamd64:
	@GOOS=linux GOARCH=amd64 go build -o bin/modversion modversion.go
	@cd bin;tar -czf modversion_v${VERSION}_linux_amd64.tar.gz modversion
	@rm -f modversion
	
.PHONY: build_darwin
build_darwin: darwinamd64
	
.PHONY: darwinamd64
darwinamd64:
	@GOOS=darwin GOARCH=amd64 go build -o bin/modversion modversion.go
	@cd bin;tar -czf modversion_v${VERSION}_darwin_amd64.tar.gz modversion
	@rm -f modversion
	
.PHONY: build_windows
build_windows: win386 winamd64

.PHONY: win386
win386:
	@GOOS=windows GOARCH=386 go build -o bin/modversion modversion.go
	@cd bin;tar -czf modversion_v${VERSION}_windows_i386.tar.gz modversion
	@rm -f modversion
	
.PHONY: winamd64
winamd64:
	@GOOS=windows GOARCH=amd64 go build -o bin/modversion modversion.go
	@cd bin;tar -czf modversion_v${VERSION}_windows_amd64.tar.gz modversion
	@rm -f modversion