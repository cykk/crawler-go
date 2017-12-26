package base

import (
	"net/http"
)

type Data interface {
	Valid() bool //是否有效
}

//请求
type Request struct {
	httpReq *http.Request
	depth   uint32
}

//创建一个请求
func NewRequest(httpReq *http.Request, depth uint32) *Request {
	return &Request{httpReq: httpReq, depth: depth}
}

func GetRequest(req *http.Request) *http.Request {
	return req.httpReq
}

func GetDepth(req *Request) uint32 {
	return req.depth
}

func (req *Request) Valid() bool {
	return req.httpReq != nil && req.httpReq.URL != nil
}

//响应
type Response struct {
	httpResp *http.Response
	depth    uint32
}

//创建一个响应
func NewResponse(httpResp *http.Response, depth uint32) *Response {
	return &Response{httpResp: httpResp, depth: depth}
}

func GetResponse(resp *Response) *http.Response {
	return resp.httpResp
}

func GetDepth(resp *Response) uint32 {
	return resp.depth
}

func (resp *Response) Valid() bool {
	return resp.httpResp != nil && resp.httpResp.Body != nil
}

//条目
type Item map[string]interface{}

func (item Item) Valid() bool {
	return item != nil
}
