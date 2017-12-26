package analyzer

import (
	"base"
)

//分析器接口类型
type Analyzer interface {
	Id() uint32                                                                     //获取ID
	Analyze(respParsers []ParseResponse, resp base.Response) ([]base.Data, []error) //根据规则分析响应并返回请求和条目
}
