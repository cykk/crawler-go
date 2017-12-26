package itempipeline

import (
	"crawler-go/base"
)

//条目处理管道接口类型
type ItemPipeline interface {
	Send(item base.Item) error //发送条目
	FailFast() bool            //是否是快速失败
	SetFailFast()              //设置快速失败
	Count() []uint64           //获得已发送、已接受和已处理的条目的计数值
	ProcessingNumer() uint64   // 获取正在被处理的条目的数量。
	Summary() string           //获取摘要信息。
}
