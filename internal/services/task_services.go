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
	if s == nil || s.Store == nil {
		return store.Task{
			Title:       title,
			Description: description,
			Priority:    priority,
		}, nil
	}

	task, err := s.Store.CreateTask(title, description, priority)
	if err != nil {
		return store.Task{}, err
	}
	return task, nil
}

func (s *TaskService) GetTaskById(id int32) (store.Task, error) {
	if s == nil || s.Store == nil {
		return store.Task{Id: id}, nil
	}

	task, err := s.Store.GetTaskById(id)
	if err != nil {
		return store.Task{}, err
	}
	return task, nil
}

func (s *TaskService) ListTasks() ([]store.Task, error) {
	if s == nil || s.Store == nil {
		return []store.Task{}, nil
	}

	tasks, err := s.Store.ListTasks()
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (s *TaskService) UpdateTask(id int32, title, description string, priority int32) (store.Task, error) {
	if s == nil || s.Store == nil {
		return store.Task{
			Id:          id,
			Title:       title,
			Description: description,
			Priority:    priority,
		}, nil
	}

	task, err := s.Store.UpdateTask(id, title, description, priority)
	if err != nil {
		return store.Task{}, err
	}
	return task, nil
}

func (s *TaskService) DeleteTask(id int32) error {
	if s == nil || s.Store == nil {
		return nil
	}

	return s.Store.DeleteTask(id)
}
