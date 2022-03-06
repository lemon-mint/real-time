GOARCH=amd64 GOOS=linux go build -ldflags "-s -w" -o realtime.linux.exe .
upx --lzma -9 realtime.linux.exe
