package analyzer

import (
	"base"
	"net/http"
)

//被用于解析http响应的函数类型
type ParseResponse func(httpResp *http.Response, depth uint32) ([]base.Data, []error)
