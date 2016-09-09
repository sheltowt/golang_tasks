package handlers

import (
	"net/http"
	"github.com/sheltowt/golang_tasks/models"
)

type TaskHandler struct {
	taskModel models.TaskModel
}

func NewTaskHandler (taskModel models.TaskModel) (TaskHandler) {
	return TaskHandler{taskModel}
}

func (taskHandler TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	taskHandler.taskModel.GetTask()
}