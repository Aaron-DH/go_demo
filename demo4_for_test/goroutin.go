package main

import (
	"fmt"
	"sync"
	"time"
)

//创建一个互斥锁
//可以让程序从并发状态变成并行状态
var lock = sync.Mutex{}

//定义一个打印字符的函数
func myprint(str string) {
	//添加锁
	lock.Lock()
	for _, value := range str {
		time.Sleep(time.Microsecond * 300)
		fmt.Printf("%c", value)
	}

	//解锁
	lock.Unlock()
}

//调用者1
func person1() {
	myprint("hello")
}

//调用者2
func person2() {
	myprint("world")
}

func main() {
	//开启go程
	go person1()
	go person2()
	//保证主线程不退出,程序不结束
	for {

	}
}
