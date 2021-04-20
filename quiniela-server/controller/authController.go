package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type task struct {
	ID      int    `json:ID`
	Name    string `json:Name`
	Content string `json:Content`
}

type allTasks []task

var tasks = allTasks{
	{
		ID:      1,
		Name:    "Task one",
		Content: "Some content",
	},
}

func Register(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "WELCOME CRIS")
}

func Setup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}
