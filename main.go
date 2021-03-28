package main

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/jotagesales/wokerpool/pool"
)

//AwaitSomeSeconds awit 1 second for simulate process.
func AwaitSomeSeconds() error {
	time.Sleep(1 * time.Second)
	return nil

}

func main() {
	tasks := []*pool.Task{
		pool.NewTask(AwaitSomeSeconds),
		pool.NewTask(AwaitSomeSeconds),
		pool.NewTask(AwaitSomeSeconds),
		pool.NewTask(AwaitSomeSeconds),
	}

	fmt.Println(tasks)
	p := pool.NewPool(tasks, 3)
	p.Run()

	var numErrors int
	for _, tasks := range p.Tasks {
		if tasks.Err != nil {
			log.Error(tasks.Err)
			numErrors++
		}

		if numErrors >= 10 {
			log.Error("Too many errors on process tasks.")
			break
		}
	}

}
