package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	note "braindumpnotes.com/notes/pkg"
	env "braindumpnotes.com/notes/pkg/env"
)

type OkMessage struct {
	Data interface{} `json:"data"`
}

type ErrorMessage struct {
	Error interface{} `json:"error"`
}

func toOkMessage(data interface{}) string {
	bytes, err := json.Marshal(OkMessage{Data: data})
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func toErrorMessage(msg string) string {
	bytes, err := json.Marshal(ErrorMessage{Error: msg})
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func health(env env.Env) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		fmt.Fprint(w, toOkMessage("ok"))
	}
}

func notesCreate(w http.ResponseWriter, req *http.Request, environ env.Env) {
	var inputNote note.InputNote
	err := json.NewDecoder(req.Body).Decode(&inputNote)
	if err != nil {
		errMsg := toErrorMessage(err.Error())
		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}

	err = environ.Store.Create(inputNote)
	if err != nil {
		errMsg := toErrorMessage(err.Error())
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	data := toOkMessage("ok")

	fmt.Fprint(w, string(data))
}

func notesUpdate(w http.ResponseWriter, req *http.Request, environ env.Env) {
}

func notesDelete(w http.ResponseWriter, req *http.Request, environ env.Env) {
}

func notesList(w http.ResponseWriter, req *http.Request, environ env.Env) {
	notes, err := environ.Store.List()
	if err != nil {
		errMsg := toErrorMessage(err.Error())
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	data := toOkMessage(notes)
	fmt.Fprint(w, string(data))
}

func notesGet(w http.ResponseWriter, req *http.Request, name string, environ env.Env) {
	noteItem, err := environ.Store.Get(name)
	if err != nil {
		errMsg := toErrorMessage(err.Error())
		http.Error(w, errMsg, http.StatusNotFound)
		return
	}
	data := toOkMessage(noteItem)
	fmt.Fprint(w, string(data))
}

func notes(environ env.Env) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		switch req.Method {
		case "GET":
			{
				keys, ok := req.URL.Query()["name"]
				if ok {
					notesGet(w, req, keys[0], environ)
				} else {
					notesList(w, req, environ)
				}
			}
		case "POST":
			notesCreate(w, req, environ)
		case "PUT":
			notesUpdate(w, req, environ)
		case "DELETE":
			notesDelete(w, req, environ)
		default:
			http.Error(w, "Unknown method", http.StatusBadRequest)
		}
	}
}

func main() {
	fmt.Println("Starting App.")
	environ := env.MakeInMemEnv()

	http.HandleFunc("/_health", health(environ))
	http.HandleFunc("/notes", notes(environ))
	http.ListenAndServe(":8090", nil)
}
