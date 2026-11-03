package main

import (
	"encoding/json"
	"errors"
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

func getTodo(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	t, err := getTodoById(id)
	if err != nil {
		writeJSON(w, http.StatusNotFound, map[string]string{"message": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, t)
}

func toggleTodoStatus(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	t, err := getTodoById(id)
	if err != nil {
		writeJSON(w, http.StatusNotFound, map[string]string{"message": err.Error()})
		return
	}

	t.Completed = !t.Completed
	writeJSON(w, http.StatusOK, t)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	// find the todo with the given id and delete it from the todos list/slice
	for i, t := range todos {
		if t.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			writeJSON(w, http.StatusOK, todos)
			return
		}
	}

	writeJSON(w, http.StatusNotFound, map[string]string{"message": "todo not found"})
}

func getTodoById(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}

	return nil, errors.New("todo not found")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /todos", getTodos)
	mux.HandleFunc("GET /todos/{id}", getTodo)
	mux.HandleFunc("POST /todos", addTodo)
	mux.HandleFunc("PATCH /todos/{id}", toggleTodoStatus)
	mux.HandleFunc("DELETE /todos/{id}", deleteTodo)

	http.ListenAndServe("localhost:9090", mux)
}
