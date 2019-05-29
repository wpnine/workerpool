package worker

import (
	"fmt"
	"sync"
	"testing"
)

var worker = NewWorkerPool(10)
var count = 0
var lock sync.Mutex
const TASK_COUNT = 100000
const EXEC_COUNT = 10000
var exitChan = make(chan bool)

func BenchmarkDoTask(b *testing.B)  {
	count = 0

	for k:= 0;k<TASK_COUNT;k++{
		worker.PostTask(func() {
			var sum = 0
			for i:=0;i< EXEC_COUNT;i++{
				sum+= i
			}
			fmt.Println("sum",sum)

			lock.Lock()
			count++
			var tempIndex = count
			lock.Unlock()

			fmt.Println("执行了",tempIndex,"个任务")
			if tempIndex == TASK_COUNT{
				exitChan <- true
			}
		})

	}

	<-exitChan
}


func BenchmarkDoTask2(b *testing.B)  {
	count = 0

	for k:= 0;k<TASK_COUNT;k++{
		go func() {

			var sum = 0
			for i := 0; i < EXEC_COUNT; i++ {
				sum += i
			}
			fmt.Println("sum", sum)

			lock.Lock()
			count++
			var tempIndex = count
			lock.Unlock()

			fmt.Println("执行了", tempIndex, "个任务")
			if tempIndex == TASK_COUNT {
				exitChan <- true
			}

		}()

	}

	<-exitChan
}