package itempipeline

import (
	"crawler-go/base"
)

//被用来处理条目的函数类型
type ProcessItem func(item base.Item) (base.Item, error)
