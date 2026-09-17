package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type taskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    int32  `json:"priority"`
}

func (api *Application) handleCreateTAsk(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var payload taskRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	payload.Title = strings.TrimSpace(payload.Title)
	if payload.Title == "" {
		http.Error(w, "title is required", http.StatusBadRequest)
		return
	}

	task, err := api.TaskService.CreateTask(payload.Title, payload.Description, payload.Priority)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(task)
}

func (api *Application) handleGetTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idParam := strings.TrimSpace(r.URL.Query().Get("id"))
	if idParam == "" {
		http.Error(w, "task id is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
		http.Error(w, "invalid task id", http.StatusBadRequest)
		return
	}

	task, err := api.TaskService.GetTaskById(int32(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(task)
}

func (api *Application) handleListTasks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	tasks, err := api.TaskService.ListTasks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(tasks)
}

func (api *Application) handleUpdateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idParam := strings.TrimSpace(r.URL.Query().Get("id"))
	if idParam == "" {
		http.Error(w, "task id is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
		http.Error(w, "invalid task id", http.StatusBadRequest)
		return
	}

	var payload taskRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	task, err := api.TaskService.UpdateTask(int32(id), payload.Title, payload.Description, payload.Priority)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(task)
}

func (api *Application) handleDeleteTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idParam := strings.TrimSpace(r.URL.Query().Get("id"))
	if idParam == "" {
		http.Error(w, "task id is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
		http.Error(w, "invalid task id", http.StatusBadRequest)
		return
	}

	if err := api.TaskService.DeleteTask(int32(id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
