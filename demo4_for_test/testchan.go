package main

import (
	"fmt"
	"math/rand"
	"time"
)

//定义管道模拟缓冲区
var myChan = make(chan int, 10)

//定义生产者函数
func producter() {
	//定义随机因子
	rand.Seed(time.Now().UnixNano())
	//生成随机数
	for i := 0; i < 10; i++ {
		num := rand.Intn(100)
		fmt.Println("生产者生产了", num)
		//将生产的数据存入到管道中
		myChan <- num
	}
}

//定义函数模拟消费者
func customer() {
	//从管道中读取数据
	for i := 0; i < 10; i++ {
		num := <-myChan
		fmt.Println("----消费者消费了", num)
	}
}

func main() {
	//创建两个go程
	//多个生产者和多个消费者
	go producter()
	go producter()

	go customer()
	go customer()

	for {

	}
}
