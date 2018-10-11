package controller

import (
	sess "Go-web-test/chitchat/model/session"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func parseTemplateFiles(filenames ...string) (t *template.Template) {
	var files []string
	t = template.New("layout")
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	t = template.Must(t.ParseFiles(files...))
	return
}

func generateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(writer, "layout", data)
}

func sessionIsOk(writer http.ResponseWriter, request *http.Request) (sess.Sesion, bool) {
	cookie, err := request.Cookie("_cookie")
	if err != nil {
		fmt.Println("sessionIsOk cookie", err.Error())
		return nil, false
	}
	sessionID, err := strconv.Atoi(cookie.Value)
	if err != nil {
		fmt.Println("sessionIsOk Atoi", err.Error())
		return nil, false
	}
	session, ok := sess.GetSessionManager(sessionID)
	if ok {
		return session, true
	}
	return nil, false
}
