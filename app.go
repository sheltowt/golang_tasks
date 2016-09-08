package main

import (
	"fmt"
	"log"
	"github.com/spf13/viper"
	"gopkg.in/mgo.v2"
    "github.com/gorilla/mux"
    "reflect"
    // "github.com/sheltowt/golang_tasks/handlers"
    // "github.com/sheltowt/golang_tasks/models"
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

	session, err := mgo.Dial(url)
	c := session.DB(database).C(collection)

	err := c.Find(query).One(&result)

	fmt.Println(reflect.TypeOf(c))

	// r := mux.NewRouter()
	// r.HandleFunc("/tasks", TasksHandler)

	// http.Handle("/", r)
}