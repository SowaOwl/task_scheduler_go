package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"task_scheduler/tasks"
	"task_scheduler/workers"
)

func main() {
	workerCount := 2
	wp := workers.NewWorkerPool(workerCount)

	tasksQuery := []tasks.Task{
		tasks.NewHttpTask("https://testscheduler.free.beeceptor.com/test", ""),
		tasks.NewFileTask("file.txt", "Test Data"),
		tasks.NewCalculateTask(10, 2),
		tasks.NewHttpTask("https://testscheduler.free.beeceptor.com/test", ""),
		tasks.NewFileTask("new_file.txt", "New Test Data"),
		tasks.NewCalculateTask(123, 0),
	}

	for _, task := range tasksQuery {
		wp.AddTask(task)
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	<-signalChan
	log.Println("Graceful shutdown initiated...")
	wp.Shutdown()
}
