package controller

import (
	"github.com/chitchat/model/talkInfo"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// GET /threads/new
// Show the new thread form page
func NewThread(writer http.ResponseWriter, request *http.Request) {
	_, ok := sessionIsOk(writer, request)
	if !ok {
		http.Redirect(writer, request, "/login", 302)
	} else {
		generateHTML(writer, nil, "layout", "private.navbar", "new.thread")
	}
}

// POST /signup
// Create the user account
func CreateThread(writer http.ResponseWriter, request *http.Request) {
	session, ok := sessionIsOk(writer, request)
	if !ok {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err := request.ParseForm()
		if err != nil {
			fmt.Println("CreateThread ParseForm ", err.Error())
			panic(err)
		}

		title := request.PostFormValue("topic")
		userID, err := strconv.Atoi(session.GetSession("userID"))
		if err != nil {
			fmt.Println("CreateThread Atoi ", err.Error())
			panic(err)
		}
		talkInfo.CreateTheme(title, userID)
		http.Redirect(writer, request, "/", 302)
	}
}

// GET /thread/read
// Show the details of the thread, including the posts and the form to write a post
func ReadThread(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	id := vals.Get("id")

	session, ok := sessionIsOk(writer, request)
	if !ok {
		http.Redirect(writer, request, "/", 302)
		// generateHTML(writer, &thread, "layout", "public.navbar", "public.thread")
	} else {
		userID, err := strconv.Atoi(session.GetSession("userID"))
		if err != nil {
			fmt.Println("ReadThread Atoi userID", err.Error())
			panic(err)
		}
		thread, ok := talkInfo.GetTheme(id, userID)
		if !ok {
			fmt.Println("ReadThread GetTheme ", err.Error())
		}
		generateHTML(writer, &thread, "layout", "private.navbar", "private.thread")
	}
}

// POST /thread/post
// Create the post
func PostThread(writer http.ResponseWriter, request *http.Request) {
	session, ok := sessionIsOk(writer, request)
	if !ok {
		http.Redirect(writer, request, "/login", 302)
	} else {
		userID, err := strconv.Atoi(session.GetSession("userID"))
		if err != nil {
			fmt.Println("PostThread Atoi userID ", err.Error())
			panic(err)
		}
		err = request.ParseForm()
		if err != nil {
			fmt.Println("PostThread ParseForm ", err.Error())
			panic(err)
		}
		body := request.PostFormValue("body")
		themeID := request.PostFormValue("uuid")
		thread, ok := talkInfo.GetTheme(themeID, userID)
		if !ok {
			fmt.Println("PostThread GetTheme ", err.Error())
		}

		thread.AddInfo(talkInfo.Information{
			Time:     time.Now(),
			UserID:   userID,
			ThemeID:  themeID,
			InfoBody: body,
		})
		url := fmt.Sprint("/thread/read?id=", themeID)
		http.Redirect(writer, request, url, 302)
	}
}
