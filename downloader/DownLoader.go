package downloader

import (
	"base"
)

//网页下载器接口类型
type PageDownLoader interface {
	Id() uint32  //获取ID
	DownLoad(req *base.Request) (*base.Response error) //发送下载请求并获取响应
}
