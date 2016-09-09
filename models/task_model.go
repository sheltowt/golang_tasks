package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"		
)

type Task struct {
	Description string `json:"description,omitempty"`
	Title string `json:"title,omitempty"`
	Done string `json:"done,omitempty"`
}

type TaskModel struct {
	DB *mgo.Collection
}

func NewTaskModel(db *mgo.Collection) (TaskModel){
	return TaskModel{db}
}

func (taskModel TaskModel) GetTask() Task {
	task := Task{}
	taskModel.DB.Find(bson.M{}).One(task)
	return task
}