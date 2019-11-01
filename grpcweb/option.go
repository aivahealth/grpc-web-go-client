package grpcweb

import (
	"net/http"

	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/encoding"
	"google.golang.org/grpc/encoding/proto"
)

var (
	defaultDialOptions = dialOptions{}
	defaultCallOptions = callOptions{
		codec: encoding.GetCodec(proto.Name),
	}
)

type dialOptions struct {
	defaultCallOptions   []CallOption
	insecure             bool
	headers              http.Header
	transportCredentials credentials.TransportCredentials
}

type DialOption func(*dialOptions)

func WithHttpHeader(key, value string) DialOption {
	return func(opt *dialOptions) {
		if opt.headers == nil {
			opt.headers = make(http.Header)
		}
		list, ok := opt.headers[key]
		if !ok {
			list = make([]string, 0)
		}
		list = append(list, value)
		opt.headers[key] = list
	}
}

func WithDefaultCallOptions(opts ...CallOption) DialOption {
	return func(opt *dialOptions) {
		opt.defaultCallOptions = opts
	}
}

func WithInsecure() DialOption {
	return func(opt *dialOptions) {
		opt.insecure = true
	}
}

func WithTransportCredentials(creds credentials.TransportCredentials) DialOption {
	return func(opt *dialOptions) {
		opt.transportCredentials = creds
	}
}

type callOptions struct {
	codec encoding.Codec
}

type CallOption func(*callOptions)

func ForceCodec(codec encoding.Codec) CallOption {
	return func(opt *callOptions) {
		opt.codec = codec
	}
}
