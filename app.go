package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sheltowt/golang_tasks/handlers"
	"github.com/sheltowt/golang_tasks/models"
	"github.com/spf13/viper"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
	"strings"
)

func main() {

	api_root := "/api/v1.0"

	viper.SetConfigName("golang_config")
	viper.AddConfigPath("$HOME/etc/golang_tasks/")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	session, err := mgo.Dial(viper.GetString("database.connection_url"))
	if err != nil {
		panic(fmt.Errorf("Fatal error db connection: %s \n", err))
	}

	c := session.DB(viper.GetString("database.database")).C(viper.GetString("database.collection"))

	taskModel := models.NewTaskModel(c)
	taskHandler := handlers.NewTaskHandler(taskModel)

	r := mux.NewRouter()
	r.HandleFunc(strings.Join([]string{api_root, "/task"}, ""), taskHandler.GetTask)

	log.Fatal(http.ListenAndServe(":3000", r))
}
