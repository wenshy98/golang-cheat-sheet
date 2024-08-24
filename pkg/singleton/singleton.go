package singleton

import (
	"fmt"
	"sync"
)

type Instance struct {
}

var once sync.Once

var instance *Instance

// getInstance once.Do 实现 singleTon
func getInstance() *Instance {
	// 保证只执行一次
	once.Do(func() {
		instance = &Instance{}
		fmt.Println("实例化SingleObject")
	})
	return instance
}
