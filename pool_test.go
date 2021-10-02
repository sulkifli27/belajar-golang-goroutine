package belajargolanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			return "New"
		},
	}

	pool.Put("Sulkifli")
	pool.Put("Sul")
	pool.Put("Kifli")

	for i := 0; i < 10; i++ {
		go func(){
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Selesai")
}