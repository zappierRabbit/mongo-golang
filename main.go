package main

import (
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"net/http"
)

func main(){
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("", )
	r.POST("", )
	r.DELETE("", )
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb-community://localhost:27017")
	if err != nil{
		panic(err)
	}
	return s
}