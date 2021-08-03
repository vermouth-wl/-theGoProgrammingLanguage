package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

/* 聊天服务器服务端代码 */
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	// 不停获取连接上来的conn 对象
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

// 对外发送消息的通道
type client chan<- string

// 所有连接的客户端
var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

// broadcaster 广播器，使用局部变量clients来记录当前连接的客户端集合
func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		// 每当有广播消息从 messages 发送进来，都会循环 clients 对里面的每个 channel 发消息
		case msg := <-messages:
			// 把所有接收到的消息广播给所有客户端
			// 发送消息通道
			for cli := range clients {
				cli <- msg
			}
		// 每当有消息从 entering 里面发送过来，就生成一个新的 key - value，相当于给 clients 里面增加一个新的 client
		case cli := <-entering:
			clients[cli] = true

		// 每当有消息从 leaving 里面发送过来，就删掉这个 key - value 对，并关闭对应的 channel
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

// handleConn 函数会为每个过来处理的 conn 都创建一个新的 channel，开启一个新的 goroutine 去把发送给这个 channel 的消息写进 conn
func handleConn(conn net.Conn) {
	// 对外发送客户消息的通道
	ch := make(chan string)
	go clientWriter(conn, ch)

	// 获取连接过来的IP地址和端口号
	who := conn.RemoteAddr().String()
	// 将欢迎信息写进channel返回给客户端
	ch <- "欢迎 " + who
	// 生成一条广播消息写进messages里
	messages <- who + " 上线"
	// 把这个channel加入到客户端集合
	entering <- ch

	// 监听客户端往conn里写的数据，每扫描到一条就将这条消息发送到广播channel中
	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
		// 此处暂时忽略input.Err()中可能存在的错误
	}

	// 如果关闭了客户端，那么把队列离开写入 leaving 交给广播函数去删除这个客户端并关闭这个客户端
	leaving <- ch
	// 广播通知其他客户端该客户端已经关闭
	messages <- who + " 下线"
	// 关闭该客户端连接
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		// 注意：忽略网络层面的错误
		fmt.Fprintln(conn, msg)
	}
}
