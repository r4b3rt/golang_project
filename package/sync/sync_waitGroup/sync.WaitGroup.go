/*
		在介绍通道的时候，如果启用了多个子协程，我们是这样实现主协程等待子协程执行完毕并退出的：
		声明一个和子协程数量一致的通道数组，然后为每个子协程分配一个通道元素，在子协程执行完毕时向对应的通道发送数据；
		然后在主协程中，
		我们依次读取这些通道接收子协程发送的数据，只有所有通道都接收到数据才会退出主协程。

		---------------代码看起来是这样的------------------------------------------------------------
		   chs:=make([]chan int, 10)
			for i:=0;i<10;i++ {
				chs[i]=make(chan int)
				go add (1,i,chs[i])
			}

			for _, ch :=range chs {
				<-ch
			}
		---------------------------------------------------------------------------
		我总感觉这样的实现有点蹩脚，不够优雅，不知道你有没有同感，那有没有更好的实现呢？
		这就要引入我们今天要讨论的主题：sync 包提供的 sync.WaitGroup 类型。

		# sync.WaitGroup 类型 #

		sync.WaitGroup 类型是开箱即用的，也是并发安全的。该类型提供了以下三个方法：

		Add：WaitGroup 类型有一个计数器，默认值是0，我们可以通过 Add 方法来增加这个计数器的值，通常我们可以通过个方法来标记需要等待的子协程数量；
		Done：当某个子协程执行完毕后，可以通过 Done 方法标记已完成，该方法会将所属 WaitGroup 类型实例计数器值减一，通常可以通过 defer 语句来调用它；
		Wait：Wait 方法的作用是阻塞当前协程，直到对应 WaitGroup 类型实例的计数器值归零，如果在该方法被调用的时候，对应计数器的值已经是 0，那么它将不会做任何事情。
		至此，你可能已经看出来了，我们完全可以组合使用 sync.WaitGroup 类型提供的方法来替代之前通道中等待子协程执行完毕的实现方法，对应代码如下：


 */

package main

func main() {
  var wg sync.WaitGroup
  wg.Add(10)
  for i:=0;i<10;i++{
	  go add_num(i,1,wg.Done)
  }
  wg.Wait()
}

func add_num (a,b int, deferFunc func()){
   defer func (){
	     deferFunc()
   }()
   c:=a+b 
   fmt.Printf("%d + %d = %d\n", a, b, c)
}


