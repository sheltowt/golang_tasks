package main

import (
	"fmt"
	"github.com/spf13/viper"
	"gopkg.in/mgo.v2"
    "github.com/gorilla/mux"
    "github.com/sheltowt/golang_tasks/handlers"
    "github.com/sheltowt/golang_tasks/models"
    "net/http"
    "log"
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

	log.Println(viper.GetString("database.connection_url"))
	log.Println(viper.GetString("database.database"))
	log.Println(viper.GetString("database.collection"))

	session, err := mgo.Dial(viper.GetString("database.connection_url"))
	log.Println(err)

	c := session.DB(viper.GetString("database.database")).C(viper.GetString("database.collection"))

	task := models.Task{}

	err = c.Find(nil).One(&task)

	taskModel := models.NewTaskModel(c)
	taskHandler := handlers.NewTaskHandler(taskModel)

	r := mux.NewRouter()
	r.HandleFunc("/task", taskHandler.GetTask)

	log.Fatal(http.ListenAndServe(":3000", r))
}