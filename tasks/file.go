package tasks

import (
	"fmt"
	"os"
)

type FileTask struct {
	path string
	data string
}

func NewFileTask(path string, data string) *FileTask {
	return &FileTask{path, data}
}

func (f *FileTask) Start() error {
	file, err := os.OpenFile(f.path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(f.data + "\n")
	return err
}

func (f *FileTask) StartMsg() string {
	return fmt.Sprintf("FileTask Start by Path %s", f.path)
}

func (f *FileTask) EndMsg() string {
	return fmt.Sprintf("FileTask End by Path %s", f.path)
}
