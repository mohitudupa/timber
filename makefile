dev:
	go run github.com/mohitudupa/timber
build:
	env GOARCH="amd64" GOOS="darwin" go build -o builds/timber-mac github.com/mohitudupa/timber
	env GOARCH="amd64" GOOS="windows" go build -o builds/timber-windows.exe github.com/mohitudupa/timber
	env GOARCH="amd64" GOOS="linux" go build -o builds/timber-linux github.com/mohitudupa/timber
