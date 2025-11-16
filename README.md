1、初始化

go mod init sc_policing_api

2、安装依赖

go mod tidy

3、编译

GOOS=linux GOARCH=amd64 go build -o libu-admin main.go
