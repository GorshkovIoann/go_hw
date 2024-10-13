package server

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type InputString struct {
	InputString string `json:"inputString"`
}

type OutputString struct {
	OutputString string `json:"outputString"`
}

type Server struct {
	router *mux.Router
}

func NewServer() *Server {
	srv := &Server{
		router: mux.NewRouter(),
	}

	// Регистрация обработчиков
	srv.router.HandleFunc("/version", getVersionHandler())
	srv.router.HandleFunc("/decode", decodeHandler())
	srv.router.HandleFunc("/hard-op", hardOpHandler())

	return srv
}

func (srv *Server) Start() {
	http.ListenAndServe(":8080", srv.router)
}

func getVersionHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "1")
	}
}

func decodeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var in InputString
		if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		decoded, err := base64.StdEncoding.DecodeString(in.InputString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		resp := OutputString{OutputString: string(decoded)}
		json.NewEncoder(w).Encode(resp)
	}
}

func hardOpHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Duration(10+rand.Intn(10)) * time.Second)

	}
}
