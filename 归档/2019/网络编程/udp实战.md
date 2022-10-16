# UDP实战
和TCP不同，UDP是面向数据报的协议，它没有诸如超时重传、ACK确认、流量控制和拥塞控制这些能力，UDP是一个不可靠的通信协议，而且不保证报文的顺序，而且可能会丢包，这些异常都需要我们自己去处理。

因此UDP比起TCP更加的轻便、简单，一般会用在一些实时性要求更高，或者对丢包、时延要求不高的场景（如游戏）。

## 一个简单的UDP案例（Golang实现）
首先是UDP服务端（连接发起方）
```go
package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	ipAddr := "127.0.0.1"
	port := 12345

	ip := net.ParseIP(ipAddr)
	udpAddr := net.UDPAddr{
		IP:   ip,
		Port: port,
	}
	udpConn, err := net.ListenUDP("udp", &udpAddr)
	if err != nil {
		fmt.Printf("listen udp error: %v", err)
		return
	}

	fmt.Printf("udp server listen on %s:%d \n", ipAddr, port)

	for {
		b := make([]byte, 1024)
		n, addr, err := udpConn.ReadFrom(b)
		if err != nil {
			fmt.Printf("receive error: %v", err)
			os.Exit(0)
		}

		if n <= 0 {
			fmt.Printf("receive invalid len: %d \n", n)
			continue
		}

		fmt.Printf("receive from: %s \n", addr.String())
		data := string(b)
		fmt.Printf("receive %d bytes: %s \n", n, data)

		rtnData := fmt.Sprintf("Hi, %s!", data)
		n, err = udpConn.WriteTo([]byte(rtnData), addr)
		if err != nil {
			fmt.Printf("send error: %v \n", err)
			continue
		}
	}
}
```
该服务端通过`net.ListenUDP("udp", &udpAddr)`监听一个指定的ip和端口，之后在死循环中不断获取客户端数据`udpConn.ReadFrom(b)`，没有数据则堵塞在此处。`readFrom`会返回一个客户端的连接信息。  
对收到的数据，拼接上`hi,`后发送回客户端`udpConn.WriteTo([]byte(rtnData), addr)`。

客户端代码
```go
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	ipAddr := "127.0.0.1"
	port := 12345
	ip := net.ParseIP(ipAddr)
	udpAddr := net.UDPAddr{
		IP:   ip,
		Port: port,
	}
	udpConn, err := net.DialUDP("udp", nil, &udpAddr)
	if err != nil {
		fmt.Printf("dial udp server error: %v", err)
		os.Exit(0)
	}
	defer udpConn.Close()

	input := bufio.NewScanner(os.Stdin)
	fmt.Printf("please input:")
	for input.Scan() {
		text := input.Text()
		fmt.Printf("now sending: %s\n", text)

		n, err := udpConn.Write([]byte(text))
		if n < 0 {
			fmt.Printf("send failed: %v\n", n)
			continue
		}
		if err != nil {
			fmt.Printf("send err: %v\n", err)
			continue
		}

		fmt.Printf("send bytes: %d\n", n)

		data := make([]byte, 1024)
		n, err = udpConn.Read(data)
		if n < 0 {
			fmt.Printf("receive failed: %v\n", n)
			continue
		}
		if err != nil {
			fmt.Printf("receive err: %v\n", err)
			continue
		}
		fmt.Printf("%s\n", string(data))
	}
}
```
客户端通过`net.DialUDP("udp", nil, &udpAddr)`对udp服务端发起通信。客户端会接收用户terminal上的输入，再发送到服务端`udpConn.Write([]byte(text))`，然后接收服务端发送过来数据`udpConn.Read(data)`。

### 启动案例和分析

启动服务端
```
$ go run udpserver.go
udp server listen on 127.0.0.1:12345
receive from: 127.0.0.1:62103 
receive 2 bytes: g1 
receive from: 127.0.0.1:62103 
receive 2 bytes: g2
```

启动客户端
```
$ go run udpclient.go
please input:g1
now sending: g1
send bytes: 2
Hi, g1
g2
now sending: g2
send bytes: 2
Hi, g2
```

此时我们再断开udp服务端连接，会发现客户端没有任何异常，我们重新启动服务端后，客户端仍旧可以继续像服务端发送数据。

服务端操作示例
```
$ go run udpserver.go
udp server listen on 127.0.0.1:12345
receive from: 127.0.0.1:56717 
receive 2 bytes: g1 
receive from: 127.0.0.1:56717 
receive 2 bytes: g2
^Csignal: interrupt

$ go run udpserver.go
udp server listen on 127.0.0.1:12345
receive from: 127.0.0.1:56717 
receive 2 bytes: g3
```

客户端操作示例
```
$ go run udpclient.go
please input:g1
now sending: g1
send bytes: 2
Hi, g1
g2
now sending: g2
send bytes: 2
Hi, g2
g3
now sending: g3
send bytes: 2
Hi, g3
```

其实，无论是先启动服务端还是先启动客户端，客户端方面都不会有任何提示，会一直阻塞。

如果服务端未启动的情况下，客户端先发送数据过去，会是什么情况？

客户端操作示例
```
$ go run udpclient.go
please input:g1
now sending: g1
send bytes: 2
receive err: read udp 127.0.0.1:61428->127.0.0.1:12345: read: connection refused
```
这里程序抛出了异常，这里就有疑问了，udp不是面向数据报文的么，或者说面向无连接的，既然无连接的，又是如何知道对方连接已经断开了。按理来说，数据发送出去后，不管成功不成功udp都不会在意的。

接下来就来聊一下这个问题。

## UDP的Connected状态
在Unix/Linux的设计中，为UDP套接字提供了一个connect函数，该函数和TCP的Connect不同，它不会引起三次握手，不会引起和服务端的任何交互。

在没有调用connect函数下，调用套接字的receive（golang中为read），会一直阻塞，等待返回（或者超时）。

但是一旦在UDP套接字中调用connect函数，便会将该套接字和服务端地址和端口建立联系，通过这种关系操作系统内核便能够将收到的信息和对应的套接字建立关联。一旦服务器断开了，套接字便能收到错误信息。

现在回过头来看之前的示例代码，我们的客户端代码中，先通过`udpConn.Write([]byte(text))`往服务端发送数据，操作系统在接收到该报文欲尝试往对应的地址和端口发送时，由于对应的地址和端口不可达，一个ICMP报文会返回给操作系统，该报文中包含了目的地址和目标端口等信息。

虽然我们没有显式的调用connect函数（golang中也没有提供这样的api），但是我们最终read的时候收到了错误回复，这个错误说明了golang中帮我们调用了connect，因此操作系统通过ICMP报文中的信息找到了对应的套接字，从而将错误信息返回给它。

## Golang中UDP的Connected状态
现在我们知道了golang中会有帮我们调用connect函数，那是不是所有udp客户端都会有呢？

其实并不是的。注意到我们上面客户端是如何连接服务端的。
```go
net.DialUDP("udp", nil, &udpAddr)
```

接下来我们看下另外一个客户端的例子
```go
func main() {
	srcIPAddr := "127.0.0.1"
	srcPort := 8888
	srcIP := net.ParseIP(srcIPAddr)
	srcUDPAddr := net.UDPAddr{
		IP:   srcIP,
		Port: srcPort,
	}

	destIPAddr := "127.0.0.1"
	destPort := 12345
	destIP := net.ParseIP(destIPAddr)
	destUDPAddr := net.UDPAddr{
		IP:   destIP,
		Port: destPort,
	}

	udpConn, err := net.ListenUDP("udp", &srcUDPAddr)
	if err != nil {
		fmt.Printf("listen udp error: %v", err)
		return
	}

	input := bufio.NewScanner(os.Stdin)
	fmt.Printf("please input:")
	for input.Scan() {
		text := input.Text()
		fmt.Printf("now sending: %s\n", text)

		n, err := udpConn.WriteTo([]byte(text), &destUDPAddr)
		if n < 0 {
			fmt.Printf("send failed: %v\n", n)
			continue
		}
		if err != nil {
			fmt.Printf("send err: %v\n", err)
			continue
		}

		fmt.Printf("send bytes: %d\n", n)

		data := make([]byte, 1024)
		n, _, err = udpConn.ReadFrom(data)
		if n < 0 {
			fmt.Printf("receive failed: %v\n", n)
			continue
		}
		if err != nil {
			fmt.Printf("receive err: %v\n", err)
			continue
		}
		fmt.Printf("%s\n", string(data))
	}
}
```
同样的，我们先运行客户端
```
$ go run udpclient.go
please input:g1
now sending: g1
send bytes: 2
```

此时服务端没有启动，可看到程序一直阻塞在read处，由于我们没有设置读取的超时时间，于是客户端会一直阻塞在这里，无法继续发送数据，即使服务端已经启动。

其实在golang中，将udp client分为了两种状态，`connected`和`unconnected`
1. 通过`DialUDP`获取的`*UDPConn`是`connected`状态，明显的在调用DialUDP时候我们已经传入了目标地址和目标端口，这时内部会调用connect，从而将套接字得到绑定。
2. 通过`ListenUDP`获取的`*UDPConn`是`unconnected`状态，此时并没有传入目标地址和目标端口，因此内部也没法调用connect，也正因此我们在发送数据的时候需要提供目标地址和目标端口。

有了这个概念，在调用读取和写入的api时才能正确
1. 如果`*UDPConn`是`connected`,读写方法是`Read`和`Write`。
2. 如果`*UDPConn`是`unconnected`,读写方法是`ReadFromUDP`和`WriteToUDP`（以及`ReadFrom`和`WriteTo`)。
