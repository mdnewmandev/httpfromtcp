package response

import (
	"fmt"
	"io"

	"github.com/mdnewmandev/httpfromtcp/internal/headers"
)

type StatusCode int

const (
	StatusCode200 StatusCode = 200
	StatusCode400 StatusCode = 400
	StatusCode500 StatusCode = 500
)

func WriteStatusLine(w io.Writer, statusCode StatusCode) error {
	var statusText string
	switch statusCode {
	case StatusCode200:
		statusText = "OK"
	case StatusCode400:
		statusText = "Bad Request"
	case StatusCode500:
		statusText = "Internal Server Error"
	default:
		statusText = ""
	}

	statusLine := fmt.Sprintf("HTTP/1.1 %d %s\r\n", statusCode, statusText)
	_, err := w.Write([]byte(statusLine))
	if err != nil {
		return err
	}
	
	return nil
}

func GetDefaultHeaders(contentLen int) headers.Headers {
	h := headers.NewHeaders()
	h.Set("Content-Length", fmt.Sprintf("%d", contentLen))
	h.Set("Connection", "close")
	h.Set("Content-Type", "text/plain")

	return h
}

func WriteHeaders(w io.Writer, headers headers.Headers) error {
	for key, value := range headers {
		headerLine := fmt.Sprintf("%s: %s\r\n", key, value)
		_, err := w.Write([]byte(headerLine))
		if err != nil {
			return err
		}
	}
	_, err := w.Write([]byte("\r\n"))
	if err != nil {
		return err
	}
	return nil
}