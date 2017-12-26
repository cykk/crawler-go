package middleware

import (
	"errors"
	"fmt"
	"reflect"
)

//实体的接口类型
type Entity interface {
	Id() uint32
}

//实体池的接口类型
type Pool interface {
	Return(entity Entity) error //将一个实体放回到实体池中
	Take() (Entity, error)      //从实体池中取出一个实体
	Total() uint32              //实体池中的总容量
	Used() uint32               //已经被使用的实体数量
}

//实体池的实现类型
type myPool struct {
	total       uint32          // 池的总容量。
	etype       reflect.Type    // 池中实体的类型。
	genEntity   func() Entity   // 池中实体的生成函数。
	container   chan Entity     // 实体容器。
	idContainer map[uint32]bool // 实体ID的容器。
	mutex       sync.Mutex      // 针对实体ID容器操作的互斥锁。
}

//创建实体池
func NewPool(total uint32, entityType reflect.Type, getEntity func() Entity) (Pool, error) {
	if total == 0 {
		errMsg := fmt.Sprintf("The pool can not be initialized! (total=%d)\n", total)
		return nil, errors.New(errMsg)
	}
	size := int(total)
	container := make(chan Entity, size)
	idContatiner := make(map[uint32]bool)
	for i = 0; i < size; i++ {
		newEntity := getEntity()
		if entityType != reflect.TypeOf(newEntity) {
			errMsg :=
				fmt.Sprintf("The type of result of function genEntity() is NOT %s!\n", entityType)
			return nil, errors.New(errMsg)
		}
		container <- newEntity
		idContatiner[newEntity.Id()] = true
	}

}

//从实体池中取出一个实体
func (pool *myPool) Take() (Entity, error) {
	entity, ok := <-pool.container
	if !ok {
		return nil, errors.New("the inner contatiner is invalide!")
	}
	pool.mutex.Lock()
	defer pool.mutex.Unlock()
	pool.idContainer[entity.Id()] = false
	return entity, nil
}

//把一个实体放回到实体池中
func (pool *myPool) Return(entity Entity) error {
	if pool.etype != reflect.TypeOf(entity) {
		errMsg := fmt.Sprintf("the entity is not this type %s!\n", pool.etype)
		return errors.New(errMsg)
	}

	entityId := entity.Id()
	casValue := pool.compareAndSetForIdContainer(entityId, false, true)
	if casValue == 1 {
		container <- entity
		return nil
	} else if casValue == 0 {
		errMsg := fmt.Sprintf("the entity (id=%d) has already in the pool! \n", entityId)
		return errors.New(errMsg)
	} else {
		errMsg := fmt.Sprintf("the entity (id=%d) is illegal \n", entityId)
		return errors.New(errMsg)
	}
}

// -1 键值对不存在
// 0  操作失败
// 1  操作成功
func (pool *myPool) compareAndSetForIdContainer(entityId uint32, oldValue bool, newValue bool) int8 {
	pool.mutex.Lock()
	defer pool.mutex.Unlock()
	v, ok := pool.idContainer[entityId]
	if !ok {
		return -1
	}
	if v != oldValue {
		return 0
	}
	pool.idContainer[entityId] = newValue
	return 1
}
