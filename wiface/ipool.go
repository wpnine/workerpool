package wiface

type IWorkerPool interface {


	PostTask(f FTask)

	Release()
}


type FTask func()