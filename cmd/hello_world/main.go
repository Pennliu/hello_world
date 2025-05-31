package main

import (
	"flag"
	"fmt"
	"log"
	"time"
)

func main() {
	log.Println("Starting hello_world application...")
	greet := flag.String("greet", "Hello, World!", "custom greeting message")
	flag.Parse()
	PrintHello(*greet)
}

func PrintHello(greet string) {
	fmt.Println(greet)
	// 新增功能：打印当前时间
	fmt.Println("Current time:", time.Now().Format(time.RFC1123))
}
