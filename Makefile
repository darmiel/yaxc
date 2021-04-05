compile:
	mkdir -p bin/
	@echo "Compiling for every OS and Platform"

	@echo "🐧 Compile for Linux"
	GOOS=linux GOARCH=amd64 go build -o ./bin/yaxc-linux-amd64 ./main.go
	GOOS=linux GOARCH=386 go build -o ./bin/yaxc-linux-386 ./main.go
	GOOS=linux GOARCH=arm go build -o ./bin/yaxc-linux-arm ./main.go
	GOOS=linux GOARCH=arm64 go build -o ./bin/yaxc-linux-arm64 ./main.go

	@echo "🍏 Compile for Apple"
	GOOS=darwin GOARCH=amd64 go build -o ./bin/yaxc-darwin-amd64 ./main.go

	@echo "🪟 Compile for Windows"
	GOOS=windows GOARCH=amd64 go build -o ./bin/yaxc-windows-amd64 ./main.go
	GOOS=windows GOARCH=386 go build -o ./bin/yaxc-windows-386 ./main.go

	@echo "🐡 Compile for FreeBSD"
	GOOS=freebsd GOARCH=amd64 go build -o ./bin/yaxc-freebsd-amd64 ./main.go
	GOOS=freebsd GOARCH=386 go build -o ./bin/yaxc-freebsd-386 ./main.go
	GOOS=freebsd GOARCH=arm go build -o ./bin/yaxc-freebsd-arm ./main.go