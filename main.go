package main

import (
	"fmt"

	go_say_hello "github.com/zuzustack/go-say-hello"
)

func  main()  {
	fmt.Println(go_say_hello.Say())
	fmt.Println(go_say_hello.SayName("Budi"))
	fmt.Println(go_say_hello.SayWorld())
}