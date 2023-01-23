package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type httpServer struct {
	Activities *Activities
}

func NewHTTPServer(addr string) *http.Server {
	var a *Activities
	var err error

	a, err = NewActivities()

	if err != nil {
		log.Fatal("can not db", err)
	}

	server := &httpServer{
		Activities: a,
	}

	r := mux.NewRouter()
	r.HandleFunc("/", server.handlePost).Methods("POST")
	r.HandleFunc("/{id}", server.handleGet).Methods("GET")
	r.HandleFunc("/", server.handleList).Methods("GET")
	return &http.Server{
		Addr:    addr,
		Handler: r,
	}
}

type IDDocument struct {
	ID int `json:"id"`
}

type ActivityDocument struct {
	Activity Activity `json:"activity"`
}

func (s *httpServer) handlePost(w http.ResponseWriter, r *http.Request) {
	var req ActivityDocument
	fmt.Printf("req: %v\n", req)
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := s.Activities.Insert(req.Activity)

	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res := IDDocument{ID: id}
	json.NewEncoder(w).Encode(res)
}

func (s *httpServer) handleGet(w http.ResponseWriter, r *http.Request) {
	println("GET - ")
	vars := mux.Vars(r)
	var id string
	id, ok := vars["id"]
	if !ok {
		http.Error(w, "NOT found", http.StatusNotFound)
		return
	}
	intId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "NOT found", http.StatusNotFound)
		return
	}
	a, err := s.Activities.Retrieve(uint64(intId))

	if err != nil {
		http.Error(w, "NOT found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(a)
}

func (s *httpServer) handleList(w http.ResponseWriter, r *http.Request) {
	println("LIST - ")
	activities, err := s.Activities.List(1)
	if err != nil {
		http.Error(w, "NOT found"+err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(activities)
}
