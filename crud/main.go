package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/julienshmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main(){

	router := httprouter.New()

	//uc := controllers.NewUserController(getSession())
	router.GET("",)
}

func getSession() *mgo.Session {

	s, err := mgo.Dial("mongoURL")
	if err != nil {
		panic(err)
	}

	return s
}