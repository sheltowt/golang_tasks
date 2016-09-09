package main

import (
	"fmt"
	"log"
	"github.com/spf13/viper"
	"gopkg.in/mgo.v2"
    "github.com/gorilla/mux"
    "github.com/sheltowt/golang_tasks/handlers"
    "github.com/sheltowt/golang_tasks/models"
    "net/http"
)

type Task struct {
	Type string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
}

func main() {
	viper.SetConfigName("golang_config")
	viper.AddConfigPath("$HOME/etc/golang_tasks/")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil { 
    	panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	session, err := mgo.Dial(viper.GetString("database.connection_url"))
	c := session.DB(viper.GetString("database.database")).C("database.collection")

	taskModel := models.NewTaskModel(c)
	taskHandler := handlers.NewTaskHandler(taskModel)

	r := mux.NewRouter()
	r.HandleFunc("/task", taskHandler.GetTask)

	http.Handle("/", r)
}