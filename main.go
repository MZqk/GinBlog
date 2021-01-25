package main

import (
	"GinBlog/routers"
	"fmt"
)

func main() {
	routers.Routers()
}

func init() {
	fmt.Println("初始化完成")
}
