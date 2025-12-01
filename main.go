package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func getLinesChannel(f io.ReadCloser) <-chan string {
	out := make(chan string, 1)

	go func() {
		defer f.Close()
		defer close(out)

		currentLineContents := ""
		for {
			buffer := make([]byte, 8)
			n, err := f.Read(buffer)
			if err != nil {
				if currentLineContents != "" {
					out <- currentLineContents
					currentLineContents = ""
				}
				if errors.Is(err, io.EOF) {
					break
				}
				fmt.Printf("error: %s\n", err.Error())
				break
			}
			str := string(buffer[:n])
			parts := strings.Split(str, "\n")
			for i := 0; i < len(parts)-1; i++ {
				out <- currentLineContents + parts[i]
				currentLineContents = ""
			}
			currentLineContents += parts[len(parts)-1]
		}
	}()

	return out
}

const port = ":42069"

func main() {
	l, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Printf("error listening for TCP traffic: %s\n", err.Error())
	}
	defer l.Close()

	fmt.Printf("listening for TCP traffic on %s\n", port)
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalf("failed connection: %s", err.Error())
		}

		go func(c net.Conn) {
			fmt.Printf("accepted connection from %s", c.RemoteAddr().String())
			lines := getLinesChannel(c)
			for line := range lines {
				fmt.Printf("%s\n", line)
			}
			c.Close()
			fmt.Printf("closed connection from %s", c.RemoteAddr().String())
		}(conn)
	}
}