package workers

import (
	"log"
	"sync"
	"task_scheduler/tasks"
)

type LeastBusyWorker struct {
	id        int
	taskQueue chan tasks.Task
	wg        *sync.WaitGroup
	active    int
	mutex     sync.Mutex
}

func NewLeastBusyWorker(id int, wg *sync.WaitGroup) *LeastBusyWorker {
	return &LeastBusyWorker{
		id:        id,
		taskQueue: make(chan tasks.Task),
		wg:        wg,
	}
}

func (w *LeastBusyWorker) Start() {
	for task := range w.taskQueue {
		w.mutex.Lock()
		w.active++
		w.mutex.Unlock()

		log.Printf("Worker %d: %s", w.id, task.StartMsg())
		if err := task.Start(); err != nil {
			log.Printf("Worker %d: error: %s", w.id, err)
		}
		log.Printf("Worker %d: %s", w.id, task.EndMsg())

		w.mutex.Lock()
		w.active--
		w.mutex.Unlock()
		w.wg.Done()
	}
}
