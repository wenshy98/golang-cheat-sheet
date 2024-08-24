package singleton

import (
	"fmt"
	"sync"
)

var lock sync.Mutex

// getDoubleCheckInstance Double Check 加锁实现单例
func getDoubleCheckInstance() *Instance {
	lock.Lock()
	defer lock.Unlock()
	if instance == nil {
		instance = &Instance{}
		fmt.Println("实例化SingleObject")
	}
	return instance
}
