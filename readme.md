

## 构建win程序

```
# 构建win命令行程序

go build -o aTcpPortTrans.exe main.go
```

## 构建win后台运行程序

```
# 构建win命令行程序

go build -ldflags "-H=windowsgui" -o aTcpPortTrans_daemon.exe main.go
```

