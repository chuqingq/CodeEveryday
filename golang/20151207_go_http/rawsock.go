package main

import (
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.ListenPacket("ip4:udp", "0.0.0.0")
	if err != nil {
		log.Fatalf("ListenPacket error: %v\n", err)
	}
	defer conn.Close()

	buf := make([]byte, 1024)
	count := 0
	const MAXCOUNT = 200000
	before := time.Now()
	for {
		_, addr, err := conn.ReadFrom(buf)
		if err != nil {
			log.Printf("ReadFrom error: %v\n", err)
		}

		// if addr.String() != "192.168.23.72" {
		// 	continue
		// }

		if addr.String() == "127.0.0.1" {
			log.Printf("ReadFrom %v\n", addr)
		}
		count += 1
		if count == MAXCOUNT {
			log.Printf("recv: %f\n", float64(count)/time.Since(before).Seconds())
			count = 0
			before = time.Now()
		}
	}
}

/*
nc -u 127.0.0.1 12345
123

此时能看到有消息过来，日志打印ReadFrom: 127.0.0.1
*/
