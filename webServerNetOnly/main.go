package main

import (
	"encoding/json"
	"net/http"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Read Book", Completed: false},
	{ID: "3", Item: "Record Video", Completed: false},
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	enc := json.NewEncoder(w)
	enc.SetIndent("", "    ")
	enc.Encode(data)
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, todos)
}

func addTodo(w http.ResponseWriter, r *http.Request) {
	var newTodo todo

	if err := json.NewDecoder(r.Body).Decode(&newTodo); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"message": "invalid request body"})
		return
	}

	todos = append(todos, newTodo)
	writeJSON(w, http.StatusCreated, newTodo)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /todos", getTodos)
	mux.HandleFunc("POST /todos", addTodo)

	http.ListenAndServe("localhost:9090", mux)
}
