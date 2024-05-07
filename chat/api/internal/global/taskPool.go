package global

import "fmt"

type TaskPool struct {
	taskPool     chan []byte
	taskPoolSize int
}

func NewTaskPool(taskPoolSize int) *TaskPool {
	return &TaskPool{
		taskPool:     make(chan []byte, taskPoolSize*64),
		taskPoolSize: taskPoolSize,
	}
}
func (tp *TaskPool) Run() {
	for i := 0; i < tp.taskPoolSize; i++ {
		go func(i int) {
			fmt.Println("init task", i)
			for task := range tp.taskPool {
				// fmt.Println("worker", i, "work")
				dispatch(task)
			}
		}(i)
	}
}
func (tp *TaskPool) Add(msg []byte) {
	tp.taskPool <- msg
}
