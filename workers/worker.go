package workers

import (
	"log"
	"sync"
	"task_scheduler/tasks"
)

type WorkerPool struct {
	tasks    chan tasks.Task
	wg       sync.WaitGroup
	stopChan chan struct{}
	taskWg   sync.WaitGroup
}

func NewWorkerPool(size int) *WorkerPool {
	wp := &WorkerPool{
		tasks:    make(chan tasks.Task, 100),
		stopChan: make(chan struct{}),
	}
	for i := 0; i < size; i++ {
		wp.wg.Add(1)
		go wp.worker(i)
	}
	return wp
}

func (wp *WorkerPool) worker(id int) {
	defer wp.wg.Done()
	for {
		select {
		case task := <-wp.tasks:
			log.Printf("Worker %d: %s", id, task.StartMsg())
			if err := task.Start(); err != nil {
				log.Printf("Worker %d: error: %s", id, err)
			}
			log.Printf("Worker %d: %s", id, task.EndMsg())

			wp.taskWg.Done()

		case <-wp.stopChan:
			log.Printf("Worker %d: stopped", id)
			return
		}
	}
}

func (wp *WorkerPool) AddTask(task tasks.Task) {
	wp.taskWg.Add(1)
	wp.tasks <- task
}

func (wp *WorkerPool) Shutdown() {
	close(wp.stopChan)
	wp.taskWg.Wait()
	close(wp.tasks)
	wp.wg.Wait()
}
