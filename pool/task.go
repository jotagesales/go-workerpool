package pool

import "sync"

type Task struct {
	// Err holds an error occurred during a task
	Err error
	f   func() error
}

// Tasks encapsulates a work item that shold go in a work pool.
func NewTask(f func() error) *Task {
	return &Task{f: f}
}

// Run runs a task and doess appropriate account via a given sync.WorkGroup
func (t *Task) Run(wg *sync.WaitGroup) {
	t.Err = t.f()
	wg.Done()
}
