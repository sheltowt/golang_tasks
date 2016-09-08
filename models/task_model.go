package models

import (
	"gopkg.in/mgo.v2"	
)

type TaskModel stuct {
	DB *mgo.Collection
}

func NewTaskModel(db *mgo.Collection) {
	return TaskModel{db}
}

