package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hahahhaaha")
	fmt.Println("我修改了代码")
	fmt.Println("这是hot分支")
	time.Sleep(time.Second * 2)
	time.Sleep(time.Second)
	fmt.Println("这次是push到远程库")
	fmt.Println("这次是pull到本地库")
}
