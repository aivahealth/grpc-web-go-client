package transport

import "net/http"

type ConnectOptions struct{
	Insecure bool
	Headers http.Header
}
