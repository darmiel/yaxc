# default
compileAll: compileDarwin compileLinux compileFreeBSD compileWindows

# dir
mkdirs: clean
	mkdir -p bin/

.PHONY: clean
clean:
	rm -r bin/
#

# compile
compileDarwin: mkdirs
	@echo ""
	@echo "🍏 Compile for Darwin"
	@echo ""
	GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o ./bin/yaxc-darwin-amd64 ./main.go

compileLinux: mkdirs
	@echo ""
	@echo "🐧 Compile for Linux"
	@echo ""
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./bin/yaxc-linux-amd64 ./main.go
	GOOS=linux GOARCH=386 go build -ldflags="-s -w" -o ./bin/yaxc-linux-386 ./main.go
	GOOS=linux GOARCH=arm go build -ldflags="-s -w" -o ./bin/yaxc-linux-arm ./main.go
	GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o ./bin/yaxc-linux-arm64 ./main.go

compileFreeBSD: mkdirs
	@echo ""
	@echo "🐡 Compile for FreeBSD"
	@echo ""
	GOOS=freebsd GOARCH=amd64 go build -ldflags="-s -w" -o ./bin/yaxc-freebsd-amd64 ./main.go
	GOOS=freebsd GOARCH=386 go build -ldflags="-s -w" -o ./bin/yaxc-freebsd-386 ./main.go
	GOOS=freebsd GOARCH=arm go build -ldflags="-s -w" -o ./bin/yaxc-freebsd-arm ./main.go

compileWindows: mkdirs
	@echo ""
	@echo "🪟 Compile for Windows"
	@echo ""
	GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o ./bin/yaxc-windows-amd64.exe ./main.go
	GOOS=windows GOARCH=386 go build -ldflags="-s -w" -o ./bin/yaxc-windows-386.exe ./main.go

# compress
compileAndCompress: compileAll
	@echo ""
	@echo "📦 Compress binaries"
	@echo ""
	upx -1 -k bin/yaxc-linux-*
	upx -1 -k bin/yaxc-darwin-*
	upx -1 -k bin/yaxc-windows-*
