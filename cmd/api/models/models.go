package models

import "io"

type Requester interface {
	GET() io.ReadCloser
	POST() io.ReadCloser
	PUT() io.ReadCloser
	DELETE() io.ReadCloser
}

type Request struct {
	URL       string
	Method    string
	ReqBody   []byte
	RespBody  io.ReadCloser
	RespRead  any
	Auth      *string
	Actions   []func()
	Multipart string
}
