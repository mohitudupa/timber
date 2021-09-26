dev:
	go run github.com/mohitudupa/timber
build:
	env GOARCH="amd64" GOOS="darwin" go build -o builds/mac/timber github.com/mohitudupa/timber
	env GOARCH="amd64" GOOS="windows" go build -o builds/windows/timber.exe github.com/mohitudupa/timber
	env GOARCH="amd64" GOOS="linux" go build -o builds/linux/timber github.com/mohitudupa/timber
