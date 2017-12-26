package middleware

import (
	"crawler-go/base"
	"errors"
	"fmt"
	"sync"
)

// 被用来表示通道管理器的状态的类型。
type ChannelManagerStatus uint8

const (
	CHANNEL_MANAGER_STATUS_UNINITIALIZED ChannelManagerStatus = 0 // 未初始化状态。
	CHANNEL_MANAGER_STATUS_INITIALIZED   ChannelManagerStatus = 1 // 已初始化状态。
	CHANNEL_MANAGER_STATUS_CLOSED        ChannelManagerStatus = 2 // 已关闭状态。
)

// 表示状态代码与状态名称之间的映射关系的字典。
var statusNameMap = map[ChannelManagerStatus]string{
	CHANNEL_MANAGER_STATUS_UNINITIALIZED: "uninitialized",
	CHANNEL_MANAGER_STATUS_INITIALIZED:   "initialized",
	CHANNEL_MANAGER_STATUS_CLOSED:        "closed",
}

// 通道管理器的接口类型。
type ChannelManager interface {
	// 初始化通道管理器。
	// 参数channelArgs代表通道参数的容器。
	// 参数reset指明是否重新初始化通道管理器。
	Init(channelArgs base.ChannelArgs, reset bool) bool
	// 关闭通道管理器。
	Close() bool
	// 获取请求传输通道。
	ReqChan() (chan base.Request, error)
	// 获取响应传输通道。
	RespChan() (chan base.Response, error)
	// 获取条目传输通道。
	ItemChan() (chan base.Item, error)
	// 获取错误传输通道。
	ErrorChan() (chan error, error)
	// 获取通道管理器的状态。
	Status() ChannelManagerStatus
	// 获取摘要信息。
	Summary() string
}

//通道管理器的实现类型
type myChannelManager struct {
	channelArgs base.ChannelArgs     //通道参数容器
	reqCh       chan base.Request    //请求通道
	respCh      chan base.Response   //响应通道
	itemCh      chan base.Item       //条目通道
	errorCh     chan base.Error      //错误通道
	status      ChannelManagerStatus //状态
	rwmutex     sync.RWMutex         //读写锁
}

//创建一个通道管理器
func NewChannelManager(channelArgs base.ChannelArgs) ChannelManager {
	chanman := &myChannelManager{}
	chanman.Init(channelArgs, true)
	return chanman
}

//初始化通道管理器
func (chanman *myChannelManager) Init(channelArgs base.ChannelArgs, reset bool) bool {
	if err := channelArgs.Check(); err != nil {
		panic(err)
	}
	if chanman.status == CHANNEL_MANAGER_STATUS_INITIALIZED && !reset {
		return false
	}

	chanman.rwmutex.Lock()
	defer chanman.rwmutex.Unlock()

	chanman.channelArgs = channelArgs
	chanman.reqCh = make(chan base.Request, channelArgs.ReqChanLen())
	chanman.respCh = make(chan base.Response, channelArgs.RespChanLen())
	channelArgs.itemCh = make(chan base.Item, channelArgs.ItemChanLen())
	channelArgs.errorCh = make(chan error, channelArgs.ErrorChanLen())
	chanman.status = CHANNEL_MANAGER_STATUS_INITIALIZED
	return true
}

//关闭通道管理器
func (chanman *myChannelManager) Close() bool {
	chanman.rwmutex.Lock()
	defer chanman.rwmutex.Unlock()
	if chanman.status == CHANNEL_MANAGER_STATUS_INITIALIZED {
		return false
	}
	close(chanman.reqCh)
	close(chanman.respCh)
	close(chanman.itemCh)
	close(chanman.errorCh)
	chanman.status = CHANNEL_MANAGER_STATUS_CLOSED
}

//获取请求通道
func (chanman *myChannelManager) ReqChan() (chan base.Request, error) {
	chanman.rwmutex.RLock()
	defer chanman.rwmutex.RUnlock()
	if err := chanman.Check(); err != nil {
		return nil, err
	}
	return chanman.reqCh, nil
}

//获取响应通道
func (chanman *myChannelManager) RespChan() (chan base.Response, error) {
	chanman.rwmutex.RLock()
	defer chanman.rwmutex.RUnlock()
	if err := chanman.Check(); err != nil {
		return nil, error
	}
	return chanman.respCh, nil
}

//获取条目通道
func (chanman *myChannelManager) ItemChan() (chan base.Item, error) {
	chanman.rwmutex.RLock()
	defer chanman.rwmutex.RUnlock()
	if err := chanman.Check(); err != nil {
		return nil, error
	}
	return chanman.itemCh, nil
}

//获取错误通道
func (chanman *myChannelManager) ErrorChan() (chan error, error) {
	chanman.rwmutex.RLock()
	defer chanman.rwmutex.RUnlock()
	if err := chanman.Check(); err != nil {
		return nil, error
	}
	return chanman.errorCh, nil
}

//检查通道管理器的状态，如果不是已初始化状态，则返回一个非nil的错误信息
func (chanman *myChannelManager) Check() error {
	if chanman.status == CHANNEL_MANAGER_STATUS_INITIALIZED {
		return nil
	}
	statusName, ok := statusNameMap[chanman.status]
	if !ok {
		statusName = fmt.Sprintf("%d", chanman.status)
	}
	errMsg := fmt.Sprintf("The undesirable status of channel manager: %s!\n", statusName)
	return errors.New(errMsg)
}

//获取通道管理器的状态
func (chanman *myChannelManager) Status() ChannelManagerStatus {
	return chanman.status
}

//通道管理器摘要模板
var chanmanSummaryTemplate = "status: %s, " +
	"requestChannel: %d/%d, " +
	"responseChannel: %d/%d, " +
	"itemChannel: %d/%d, " +
	"errorChannel: %d/%d"

//获取通道管理器的摘要信息
func (chanman *myChannelManager) Summary() string {
	summary := fmt.Sprintf(chanmanSummaryTemplate, statusNameMap[chanman.status],
		len(chanman.reqCh), cap(chanman.reqCh),
		len(chanman.respCh), cap(chanman.respCh),
		len(chanman.itemCh), cap(chanman.itemCh),
		len(chanman.errorCh), cap(chanman.errorCh))
	return summary
}
