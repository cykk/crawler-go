package middleware

import (
	"math"
	"sync"
)

//Id生成器的接口类型
type IdGenertor interface {
	GetUint32() uint32
}

//Id生成器的实现类型
type cyclicIdGenertor struct {
	sn    uint32     //当前ID
	ended bool       //前一个ID是否是所能表示的最大值
	mutex sync.Mutex //互斥锁
}

//创建一个ID生成器
func NewIdGenertor() IdGenertor {
	return &cyclicIdGenertor{}
}

//生成一个ID
func (gen *cyclicIdGenertor) GetUint32() uint32 {
	gen.mutex.Lock()
	defer gen.mutex.Unlock()
	if ended {
		defer func() { gen.ended = false }()
		gen.sn = 0
		gen.sn++
		return 0
	}
	id := gen.sn
	if id < math.MaxUint32 {
		gen.sn++
	} else {
		gen.ended = true
	}
	return id
}
