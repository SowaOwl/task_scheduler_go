package tasks

type Task interface {
	Start() error
	StartMsg() string
	EndMsg() string
}
