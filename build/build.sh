GOOS=darwin GOARCH=arm64 go build -o bin/ssh-http-proxy-darwin-arm64 ..
GOOS=darwin GOARCH=amd64 go build -o bin/ssh-http-proxy-darwin-amd64 ..
GOOS=linux GOARCH=amd64 go build -o bin/ssh-http-proxy-linux-amd64 ..
GOOS=windows GOARCH=amd64 go build -o bin/ssh-http-proxy-windows-amd64.exe ..
