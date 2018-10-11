package controller

import (
	sess "Go-web-test/chitchat/model/session"
	"Go-web-test/chitchat/model/sql"
	"Go-web-test/chitchat/model/user"
	"fmt"
	"net/http"
	"strconv"
)

// GET /login
// Show the login page
func Login(writer http.ResponseWriter, request *http.Request) {
	t := parseTemplateFiles("login.layout", "public.navbar", "login")
	t.Execute(writer, nil)
}

// GET /signup
// Show the signup page
func Signup(writer http.ResponseWriter, request *http.Request) {
	generateHTML(writer, nil, "login.layout", "public.navbar", "signup")
}

// POST /signup
// Create the user account
func SignupAccount(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		fmt.Println("SignupAccount ParseForm ", err.Error())
		panic(err)
	}
	normalUser := user.User{
		UserName: request.PostFormValue("name"),
		Email:    request.PostFormValue("email"),
		PassWord: sql.Encrypt(request.PostFormValue("password")),
	}
	if _, err := user.CreateUser(&normalUser); err != nil {
		fmt.Println("SignupAccount CreateUser ", err.Error())
		panic(err)
	}
	http.Redirect(writer, request, "/login", 302)
}

// POST /authenticate
// Authenticate the user given the email and password
func Authenticate(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	normalUser, err := user.SelectUser(request.PostFormValue("email"))
	if err != nil {
		fmt.Println("Authenticate SelectUser ", err.Error())
		panic(err)
	}

	if normalUser.PassWord == sql.Encrypt(request.PostFormValue("password")) {
		_ = sess.Init(normalUser.UserID)
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    strconv.Itoa(normalUser.UserID),
			HttpOnly: true,
		}
		http.SetCookie(writer, &cookie)
		http.Redirect(writer, request, "/", 302)
	} else {
		http.Redirect(writer, request, "/login", 302)
	}

}

// GET /logout
// Logs the user out
func Logout(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("_cookie")
	if err != http.ErrNoCookie {
		index, err := strconv.Atoi(cookie.Value)
		if err != nil {
			fmt.Println("Logout Atoi ", err.Error())
			panic(err)
		}
		sess.RemoveSessionManager(index)
	}
	http.Redirect(writer, request, "/", 302)
}
