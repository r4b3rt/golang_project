//time.After()特性可以配合select实现计时器
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	// 并发执行
	go func() {
		ch1 <- 1
	}()

	for {
		select {
		case e1 := <-ch1:
			//如果ch1通道成功读取数据，则执行该case处理语句
			fmt.Printf("1th case is selected. e1=%v\n", e1)
		case <-time.After(time.Second * 2):
			fmt.Println("Timed out")
		}
	}
}
