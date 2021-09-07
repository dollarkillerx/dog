build_linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dog cmd/cli.go

build_win:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dog cmd/cli.go
