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

type BusyWorkerPool struct {
	workers []*LeastBusyWorker
	wg      *sync.WaitGroup
}

func NewBusyWorkerPool(workers int) *BusyWorkerPool {
	wp := &BusyWorkerPool{
		wg: &sync.WaitGroup{},
	}
	for i := 0; i < workers; i++ {
		worker := NewLeastBusyWorker(i, wp.wg)
		wp.workers = append(wp.workers, worker)
		go worker.Start()
	}
	return wp
}

func (wp *BusyWorkerPool) AddTask(task tasks.Task) {
	leastBusyWorker := wp.workers[0]
	for _, worker := range wp.workers {
		worker.mutex.Lock()
		if worker.active < leastBusyWorker.active {
			leastBusyWorker = worker
		}
		worker.mutex.Unlock()
	}

	wp.wg.Add(1)
	leastBusyWorker.taskQueue <- task
}

func (wp *BusyWorkerPool) Shutdown() {
	wp.wg.Wait()
	for _, worker := range wp.workers {
		close(worker.taskQueue)
	}
}
