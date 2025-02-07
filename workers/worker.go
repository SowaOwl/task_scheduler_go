package workers

import "task_scheduler/tasks"

type WorkerPool interface {
	AddTask(task tasks.Task)
	Shutdown()
}
