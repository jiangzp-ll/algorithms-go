package main

import "github.com/zepeng-jiang/go-basic-demo/main/queue"

func main() {
	q := queue.NewCircularQueue(5)
	for i := 0; i < 5; i++ {
		_ = q.EnQueue(i)
	}
}
