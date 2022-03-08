GOARCH=amd64 GOOS=linux go build -ldflags "-s -w" -o realtime.linux.exe .
upx --lzma -v --best realtime.linux.exe
