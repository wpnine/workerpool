package worker

import (
	"GoTest/worker/wiface"
	"sync"
)

type WorkerPool struct {
	wiface.IWorkerPool

	lock         sync.Mutex
	workers      []chan wiface.FTask
	executeCount int
	maxPoolCount int
}



func NewWorkerPool(maxPoolSize int) wiface.IWorkerPool {
	worker := &WorkerPool{
		workers:      make([]chan wiface.FTask,maxPoolSize),
		maxPoolCount: maxPoolSize,
	}

	worker.startWorkPool()

	return worker
}

func (this *WorkerPool) startWorkPool() {
	for index,_:= range this.workers{
		this.workers[index] = make(chan wiface.FTask,4096)
		go this.startOneWorker(index,this.workers[index])
	}
}


func (this *WorkerPool) startOneWorker(index int,channel chan wiface.FTask)  {
	for{
		select {
		case task := <-channel:
			task()
		}
	}

}

func (this *WorkerPool)Release() {
	for _,worker := range this.workers{
		close(worker)
	}
}


func (this *WorkerPool)GetNextWorkerId() int {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.executeCount += 1

	return this.executeCount % this.maxPoolCount
}


func (this *WorkerPool) PostTask(task wiface.FTask) {
	this.workers[this.GetNextWorkerId()] <- task
}








