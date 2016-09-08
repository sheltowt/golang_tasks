package handlers

import (
	"github.com/sheltowt/golang_tasks/models"
)

type TaskHandler stuct {
	taskModel models.TaskModel
}

func NewTaskHandler (taskModel models.TaskModel) {
	return NewTaskHandler{taskModel}
}

