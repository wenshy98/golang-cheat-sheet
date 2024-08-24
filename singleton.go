package singleton

import (
	"fmt"
	"sync"
)

type singleObject struct {
}

var once sync.Once

var instance *singleObject

// getInstance once.Do 实现 singleTon
func getInstance() *singleObject {
	// 保证只执行一次
	once.Do(func() {
		instance = &singleObject{}
		fmt.Println("实例化SingleObject")
	})
	return instance
}

var lock sync.Mutex

// getDoubleCheckInstance Double Check 加锁实现单例
func getDoubleCheckInstance() *singleObject {
	lock.Lock()
	defer lock.Unlock()
	if instance == nil {
		instance = &singleObject{}
		fmt.Println("实例化SingleObject")
	}
	return instance
}
