package models

import (
	"log"
	"gopkg.in/mgo.v2"	
)

type Task struct {
	Description string `json:"description,omitempty"`
	Title string `json:"title,omitempty"`
	Done int `json:"done,omitempty"`
}

type TaskModel struct {
	DB *mgo.Collection
}

func NewTaskModel(db *mgo.Collection) (TaskModel){
	return TaskModel{db}
}

func (taskModel TaskModel) GetTask() (Task) {
	task := Task{}

	err := taskModel.DB.Find(nil).One(&task)
	if err != nil {
		log.Println(err.Error())
	}

	return task
}