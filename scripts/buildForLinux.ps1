$env:GOOS="linux"
$env:GOARCH="amd64"
$env:CGO_ENABLED=0
go build -o ./artifact/accessRole .