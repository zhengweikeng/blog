package main

import (
	"fmt"
	"sync"
	"time"
)

type Config struct {
	AppName string
}

var config Config

var once sync.Once

func CreateSingleInstance() {
	once.Do(func() {
		fmt.Println("CreateSingleInstance")
		config = Config{
			AppName: "test",
		}
	})
}

func main() {
	for i := 0; i < 2; i++ {
		go CreateSingleInstance()
	}

	time.Sleep(1 * time.Second)
	fmt.Println(config.AppName)
}
