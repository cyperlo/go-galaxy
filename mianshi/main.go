package main

// 现在有一个题 输入路径 /A /A/B
import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	for {
		select {
		case x := <-ch1:
			fmt.Println(x)
		case y := <-ch2:
			fmt.Println(y)
		case <-time.After(5 * time.Second):
			fmt.Println("超时退出")
			return
		}
	}

	fmt.Println("end")
}
