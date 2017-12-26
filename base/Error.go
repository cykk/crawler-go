package base

import (
	"bytes"
	"fmt"
)

//错误类型
type ErrorType string

// 错误类型常量
const (
	DOWNLOADER_ERROR     ErrorType = "Downloader Error"
	ANALYZER_ERROR       ErrorType = "Analyzer Error"
	ITEM_PROCESSOR_ERROR ErrorType = "Item Processor Error"
)

//爬虫错误接口
type CrawlerError interface {
	Type() ErrorType //获取错误类型
	Error() string   //获取错误信息
}

//爬虫错误
type myCrawlerError struct {
	errType    ErrorType //错误类型
	errMsg     string    //错误信息
	fullErrMsg string    //完成的错误信息
}

//新建一个爬虫错误
func NewCrawlerError(errType ErrorType, errMsg string) *myCrawlerError {
	return &myCrawlerError{errType: errType, errMsg: errMsg}
}

func (err *myCrawlerError) Type() ErrorType {
	return err.errType
}

func (err *myCrawlerError) Error() string {
	if err.fullErrMsg == "" {
		err.genFullErrMsg()
	}
	return err.fullErrMsg
}

//拼接完整的错误信息
func (err *myCrawlerError) genFullErrMsg() {
	var buffer bytes.Buffer
	buffer.WriteString("Crawler Error:")
	if err.errType != "" {
		buffer.WriteString(string(err.errType))
		buffer.WriteString(":")
	}
	buffer.WriteString(err.errMsg)
	err.fullErrMsg = fmt.Sprintf("%\n", buffer.String())
}
