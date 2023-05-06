package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
	"strconv"
)

func main() {
	// 命令行参数解析
	localPort := flag.Int("l", 0, "local port to listen")
	remoteAddr := flag.String("r", "", "remote address to forward")
	flag.Parse()
	// 输入提示
	if *localPort == 0 || *remoteAddr == "" {
		flag.Usage()
		os.Exit(1)
	}
	// 监听本地端口
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(*localPort))
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	// 接收客户端连接并转发
	log.Printf("listening on %d, forwarding to %s\n", *localPort, *remoteAddr)
	for {
		clientConn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go forward(clientConn, *remoteAddr)
	}
}

// 转发两个连接之间的数据
func forward(clientConn net.Conn, remoteAddr string) {
	// 连接远程地址
	remoteConn, err := net.Dial("tcp", remoteAddr)
	if err != nil {
		log.Println(err)
		return
	}
	//defer remoteConn.Close()
	// 转发数据
	//log.Printf("forwarding from %s to %s\n", clientConn.RemoteAddr(), remoteConn.RemoteAddr())
	go func() {
		defer clientConn.Close()
		_, err := io.Copy(remoteConn, clientConn)
		if err != nil {
			log.Println("failed to copy from connection:", err)
		}
	}()
	go func() {
		defer remoteConn.Close()
		_, err := io.Copy(clientConn, remoteConn)
		if err != nil {
			log.Println("failed to copy from connection:", err)
		}
	}()
}
