package interfc

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	note "braindumpnotes.com/notes/pkg"
	env "braindumpnotes.com/notes/pkg/env"
)

type OkMessage struct {
	Data interface{} `json:"data"`
}

type ErrorMessage struct {
	Error interface{} `json:"error"`
}

func enableCors(w *http.ResponseWriter) {
	// https://flaviocopes.com/golang-enable-cors/
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, PATCH, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
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

func notesCreate(w http.ResponseWriter, req *http.Request, envi env.Env) {
	var inputNote note.InputNote
	err := json.NewDecoder(req.Body).Decode(&inputNote)
	if err != nil {
		errMsg := toErrorMessage(err.Error())
		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}

	err = envi.NoteStore.Create(inputNote)
	if err != nil {
		errMsg := toErrorMessage(err.Error())
		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}

	data := toOkMessage("ok")
	fmt.Fprint(w, string(data))
}

func notesUpdate(w http.ResponseWriter, req *http.Request, fname string, envi env.Env) {
	var updateNote note.UpdateNote
	err := json.NewDecoder(req.Body).Decode(&updateNote)
	if err != nil {
		errMsg := toErrorMessage(err.Error())
		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}

	_, err = envi.NoteStore.Get(fname)
	if err != nil {
		errMsg := toErrorMessage(err.Error())
		http.Error(w, errMsg, http.StatusNotFound)
		return
	}

	err = envi.NoteStore.Update(fname, updateNote)
	if err != nil {
		errMsg := toErrorMessage(err.Error())
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	data := toOkMessage("ok")
	fmt.Fprint(w, string(data))
}

func notesDelete(w http.ResponseWriter, req *http.Request, envi env.Env) {
}

func notesList(w http.ResponseWriter, req *http.Request, acctId string, envi env.Env) {
	notes, err := envi.NoteStore.List()
	if err != nil {
		errMsg := toErrorMessage(err.Error())
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	data := toOkMessage(notes)
	fmt.Fprint(w, string(data))
}

func notesGet(w http.ResponseWriter, req *http.Request, acctId string, name string, envi env.Env) {
	noteItem, err := envi.NoteStore.Get(name)
	if err != nil {
		errMsg := toErrorMessage(err.Error())
		http.Error(w, errMsg, http.StatusNotFound)
		return
	}
	data := toOkMessage(noteItem)
	fmt.Fprint(w, string(data))
}

func notes(envi env.Env) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		enableCors(&w)

		switch req.Method {
		case "OPTIONS": return
		case "GET":
			{
				acctId, ok := req.URL.Query()["accountId"]
				if !ok {
					errMsg := toErrorMessage("Account Id required in query parameters")
					http.Error(w, errMsg, http.StatusBadRequest)
					return
				}
				name, ok := req.URL.Query()["name"]
				if ok {
					notesGet(w, req, acctId[0], name[0], envi)
				} else {
					notesList(w, req, acctId[0], envi)
				}
			}
		case "POST":
			notesCreate(w, req, envi)
		case "PATCH":
			{
				uriSegments := strings.Split(
					strings.TrimSuffix(req.URL.Path, "/"),
					"/")

				fname := uriSegments[len(uriSegments)-1]
				notesUpdate(w, req, fname, envi)
			}
		case "DELETE":
			notesDelete(w, req, envi)
		default:
			http.Error(w, "Unknown method", http.StatusBadRequest)
		}
	}
}

func accountGet(w http.ResponseWriter, req *http.Request, email note.Email, env env.Env) {
	accountItem, err := env.AccountStore.Get(email)
	if err != nil {
		errMsg := toErrorMessage(err.Error())
		http.Error(w, errMsg, http.StatusNotFound)
		return
	}
	data := toOkMessage(accountItem)
	fmt.Fprint(w, string(data))
}

func accounts(envi env.Env) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		enableCors(&w)

		switch req.Method {
		case "GET":
			{
				keys, ok := req.URL.Query()["email"]
				if ok {
					email := note.Email(keys[0])
					accountGet(w, req, email, envi)
				} else {
					errMsg := toErrorMessage("Listing accounts is not supported")
					http.Error(w, errMsg, http.StatusBadRequest)
					return
				}
			}
		case "POST":
		case "PUT":
		case "DELETE":
		default:
			http.Error(w, "Unknown method", http.StatusBadRequest)
			return
		}
	}
}

/*
Examples:
https://gist.github.com/reagent/043da4661d2984e9ecb1ccb5343bf438
*/
func Run(envi env.Env) {
	fmt.Println("Starting App.")

	http.HandleFunc("/_health", health(envi))
	http.HandleFunc("/notes", notes(envi)) // require account ID
	http.HandleFunc("/notes/", notes(envi)) // require account ID
	http.HandleFunc("/accounts", accounts(envi))
	http.ListenAndServe(":8090", nil)
}
