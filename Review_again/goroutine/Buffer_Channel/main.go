package main

import "fmt"

func main() {
	c := make(chan int, 1) //修改2为1就报错，修改2为3可以正常运行
	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Println(<-c)
}

//修改为1报如下的错误:
//fatal error: all goroutines are asleep - deadlock!

//如果想使用1作为缓冲你必须接收数据

/*
package main

import "fmt"

func main() {
    c := make(chan int, 1)

    go func() {
        fmt.Println(<-c) // 在另一个 Goroutine 中接收
        fmt.Println(<-c)
    }()

    c <- 1
    c <- 2

    // 确保主 Goroutine等待子 Goroutine完成
    // 使用 `sync.WaitGroup` 更规范
}

*/
