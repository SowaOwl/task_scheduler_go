package tasks

import (
	"errors"
	"fmt"
	"log"
)

type CalculateTask struct {
	a int
	b int
}

func NewCalculateTask(a int, b int) *CalculateTask {
	return &CalculateTask{a: a, b: b}
}

func (c *CalculateTask) Start() error {
	if c.b <= 0 {
		return errors.New("b must be greater than zero")
	}

	log.Printf("calculating %d/%d", c.a, c.b)
	log.Printf("result:  %d", c.a/c.b)

	return nil
}

func (c *CalculateTask) StartMsg() string {
	return fmt.Sprintf("Calculate Start %d/%d", c.a, c.b)
}

func (c *CalculateTask) EndMsg() string {
	return fmt.Sprintf("Calculate End %d/%d", c.a, c.b)
}
