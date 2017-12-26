package base

import (
	"errors"
)

//参数容器的接口
type Args interface {
	Check() error   //自检参数的有效性
	String() string // 参数容器的字符串表示形式
}

//通道参数容器类型
type ChannelArgs struct {
	reqChanLen   uint   //请求通道的长度
	respChanLen  uint   //响应通道的长度
	itemChanLen  uint   //条目通道的长度
	errorChanLen uint   //错误通道的长度
	description  string //描述
}

// 通道参数的容器的描述模板。
var channelArgsTemplate string = "{ reqChanLen: %d, respChanLen: %d," +
	" itemChanLen: %d, errorChanLen: %d }"

//创建通道参数的容器
func NewChannelArgs(reqChanLen uint, respChanLen uint, itemChanLen uint, errorChanLen uint) ChannelArgs {
	return ChannelArgs{reqChanLen: reqChanLen, respChanLen: respChanLen, itemChanLen: itemChanLen, errorChanLen: errorChanLen}
}

func (args *ChannelArgs) Check() error {
	if args.reqChanLen == 0 {
		return errors.New("The request channel max length (capacity) can not be 0!\n")
	}
	if args.respChanLen == 0 {
		return errors.New("The response channel max length (capacity) can not be 0!\n")
	}
	if args.itemChanLen == 0 {
		return errors.New("The item channel max length (capacity) can not be 0!\n")
	}
	if args.errorChanLen == 0 {
		return errors.New("The error channel max length (capacity) can not be 0!\n")
	}
	return nil
}

func (args *ChannelArgs) String() string {
	if args.description == "" {
		args.description =
			fmt.Sprintf(channelArgsTemplate,
				args.reqChanLen,
				args.respChanLen,
				args.itemChanLen,
				args.errorChanLen)
	}
	return args.description
}

// 获得请求通道的长度。
func (args *ChannelArgs) ReqChanLen() uint {
	return args.reqChanLen
}

// 获得响应通道的长度。
func (args *ChannelArgs) RespChanLen() uint {
	return args.respChanLen
}

// 获得条目通道的长度。
func (args *ChannelArgs) ItemChanLen() uint {
	return args.itemChanLen
}

// 获得错误通道的长度。
func (args *ChannelArgs) ErrorChanLen() uint {
	return args.errorChanLen
}

// 池基本参数容器的描述模板。
var poolBaseArgsTemplate string = "{ pageDownloaderPoolSize: %d," +
	" analyzerPoolSize: %d }"

// 池基本参数的容器。
type PoolBaseArgs struct {
	pageDownloaderPoolSize uint32 // 网页下载器池的尺寸。
	analyzerPoolSize       uint32 // 分析器池的尺寸。
	description            string // 描述。
}

// 创建池基本参数的容器。
func NewPoolBaseArgs(
	pageDownloaderPoolSize uint32,
	analyzerPoolSize uint32) PoolBaseArgs {
	return PoolBaseArgs{
		pageDownloaderPoolSize: pageDownloaderPoolSize,
		analyzerPoolSize:       analyzerPoolSize,
	}
}

func (args *PoolBaseArgs) Check() error {
	if args.pageDownloaderPoolSize == 0 {
		return errors.New("The page downloader pool size can not be 0!\n")
	}
	if args.analyzerPoolSize == 0 {
		return errors.New("The analyzer pool size can not be 0!\n")
	}
	return nil
}

func (args *PoolBaseArgs) String() string {
	if args.description == "" {
		args.description =
			fmt.Sprintf(poolBaseArgsTemplate,
				args.pageDownloaderPoolSize,
				args.analyzerPoolSize)
	}
	return args.description
}

// 获得网页下载器池的尺寸。
func (args *PoolBaseArgs) PageDownloaderPoolSize() uint32 {
	return args.pageDownloaderPoolSize
}

// 获得分析器池的尺寸。
func (args *PoolBaseArgs) AnalyzerPoolSize() uint32 {
	return args.analyzerPoolSize
}
