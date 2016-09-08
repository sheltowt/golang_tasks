package main

import (
	"fmt"
	"log"
	"github.com/spf13/viper"
	"gopkg.in/mgo.v2"
    "encoding/json"
    "log"
    "net/http"
 
    "github.com/gorilla/mux"
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

}