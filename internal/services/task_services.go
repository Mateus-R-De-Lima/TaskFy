package services

import (
	"github.com/Mateus-R-De-Lima/taskfy/internal/store"
)

type TaskService struct {
	Store store.TaskStore
}

func NewTaskService(store store.TaskStore) *TaskService {
	return &TaskService{
		Store: store,
	}
}

func (s *TaskService) CreateTask(title, description string, priority int32) (store.Task, error) {
	// Add Regras de negocios

	task, err := s.Store.CreateTask(title, description, priority)

	if err != nil {
		return store.Task{}, err
	}

	return task, err
}
