package main

import (
	"fmt"
	"sync"
)

// Singleton 结构体
type Singleton struct {
	Data int
}

var count int
var instance *Singleton
var once sync.Once

// GetInstance 获取单例实例的指针
func GetInstance() *Singleton {
	once.Do(func() { // once 只能执行一次
		count++
		instance = &Singleton{Data: count}
	})
	return instance
}

func main() {
	s1 := GetInstance()
	s2 := GetInstance()

	fmt.Println(s1 == s2) // 输出: true
	fmt.Println("s1.Data:", s1.Data)
	fmt.Println("s2.Data:", s2.Data)
	test()
}

func test() { //高并发生成100百个看看构造方法调用了几次
	var wg sync.WaitGroup
	instanceCount := 100
	instances := make([]*Singleton, instanceCount)

	// 在多个 goroutine 中获取 Singleton 实例
	for i := 0; i < instanceCount; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			instances[i] = GetInstance()
		}(i)
	}

	wg.Wait()

	for i := 1; i < instanceCount; i++ {
		if instances[i] != instances[0] {
			fmt.Printf("实例不唯一: instance[%d] != instance[0]", i)
		}
		fmt.Printf("第%d个实例的数据为%d, 构造方法用了%d次\n", i, instances[i].Data, count)
	}
}
