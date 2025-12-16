package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	dst, err := net.ResolveUDPAddr("udp", "localhost:42069")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.DialUDP("udp", nil, dst)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	stream := bufio.NewReader(os.Stdin)
	for {
		fmt.Println(">")
		line, err := stream.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		_, err = conn.Write([]byte(line))
		if err != nil {
			log.Fatal(err)
		}
	}
}