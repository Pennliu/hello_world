package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	log.Println("Starting hello_world application...")
	PrintHello()
}

func PrintHello() {
	fmt.Println("Hello, World!")
	// 新增功能：打印当前时间
	fmt.Println("Current time:", time.Now().Format(time.RFC1123))
}
