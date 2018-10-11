package main

import (
	"Go-web-test/chitchat/config"
	"Go-web-test/chitchat/controller"
	_ "Go-web-test/chitchat/model/session"
	_ "Go-web-test/chitchat/model/sql"
	_ "Go-web-test/chitchat/model/talkInfo"
	"fmt"
	"net/http"
	"time"
)

var conf config.Configuration

func main() {
	config.LoadConfig(&conf)
	fmt.Println("ChitChat 1.0 started at " + conf.Address)
	// handle static assets
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(conf.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	//
	// all route patterns matched here
	// route handler functions defined in other files
	//

	// index
	mux.HandleFunc("/", controller.Index)
	// error
	mux.HandleFunc("/err", controller.Err)

	// defined in route_auth.go
	mux.HandleFunc("/login", controller.Login)
	mux.HandleFunc("/logout", controller.Logout)
	mux.HandleFunc("/signup", controller.Signup)
	mux.HandleFunc("/signup_account", controller.SignupAccount)
	mux.HandleFunc("/authenticate", controller.Authenticate)

	// defined in route_thread.go
	mux.HandleFunc("/thread/new", controller.NewThread)
	mux.HandleFunc("/thread/create", controller.CreateThread)
	mux.HandleFunc("/thread/post", controller.PostThread)
	mux.HandleFunc("/thread/read", controller.ReadThread)

	// starting up the server
	server := &http.Server{
		Addr:           conf.Address,
		Handler:        mux,
		ReadTimeout:    time.Duration(conf.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(conf.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
