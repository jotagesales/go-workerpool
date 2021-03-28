package pool

import "sync"

//Pool is a worker group that runs a nunber of tasks at a configurated concurrency.
type Pool struct {
	Tasks       []*Task
	concurrency int
	taskChan    chan *Task
	wg          sync.WaitGroup
}

// NewPool initializes a new pool with the given tasks and at the given concurrency.
func NewPool(tasks []*Task, concurrency int) *Pool {
	return &Pool{
		Tasks:       tasks,
		concurrency: concurrency,
		taskChan:    make(chan *Task),
	}
}

//Run runs al work within the pool and block until it's finished.
func (p *Pool) Run() {
	for i := 0; i < p.concurrency; i++ {
		p.work()
	}

	p.wg.Add(len(p.Tasks))

	for _, task := range p.Tasks {
		p.taskChan <- task
	}

	close(p.taskChan)
	p.wg.Wait()

}

// work, the work loop for any single goroutine.
func (p *Pool) work() {
	for task := range p.taskChan {
		task.Run(&p.wg)
	}
}
