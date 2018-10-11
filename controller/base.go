package controller

import (
	"github.com/chitchat/model/talkInfo"
	"fmt"
	"net/http"
	"strconv"
)

// GET /err?msg=
// shows the error message page
func Err(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	_, ok := sessionIsOk(writer, request)
	if !ok {
		generateHTML(writer, vals.Get("msg"), "layout", "public.navbar", "error")
	} else {
		generateHTML(writer, vals.Get("msg"), "layout", "private.navbar", "error")
	}
}

func Index(writer http.ResponseWriter, request *http.Request) {
	session, ok := sessionIsOk(writer, request)
	if !ok {
		http.Redirect(writer, request, "/login", 302)
	} else {
		userID, err := strconv.Atoi(session.GetSession("userID"))
		if err != nil {
			fmt.Println("Index", err.Error())
			panic(err)
		}
		threads := talkInfo.GetThemeList(userID)
		generateHTML(writer, &threads, "layout", "private.navbar", "index")
	}
}
