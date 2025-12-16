package response

import (
	"fmt"

	"github.com/mdnewmandev/httpfromtcp/internal/headers"
)

type StatusCode int

const (
	StatusCode200 StatusCode = 200
	StatusCode400 StatusCode = 400
	StatusCode500 StatusCode = 500
)

func getStatusLine(statusCode StatusCode) []byte {
	reasonPhrase := ""
	switch statusCode {
	case StatusCode200:
		reasonPhrase = "OK"
	case StatusCode400:
		reasonPhrase = "Bad Request"
	case StatusCode500:
		reasonPhrase = "Internal Server Error"
	}
	return []byte(fmt.Sprintf("HTTP/1.1 %d %s\r\n", statusCode, reasonPhrase))
}

func GetDefaultHeaders(contentLen int) headers.Headers {
	h := headers.NewHeaders()
	h.Set("Content-Length", fmt.Sprintf("%d", contentLen))
	h.Set("Connection", "close")
	h.Set("Content-Type", "text/plain")

	return h
}
