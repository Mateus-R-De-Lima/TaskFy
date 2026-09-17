package pgstore

import (
	"context"

	"github.com/Mateus-R-De-Lima/taskfy/internal/store"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PGTaskStore struct {
	Queries *Queries
	Pool    *pgxpool.Pool
}

func NewPGTaskStore(pool *pgxpool.Pool) PGTaskStore {
	return PGTaskStore{
		Queries: New(pool),
		Pool:    pool,
	}
}

func mapTask(task Task) store.Task {
	return store.Task{
		Id:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Priority:    task.Priority,
		CreatedAt:   task.CreatedAt.Time,
		UpdatedAt:   task.UpdatedAt.Time,
	}
}

func (pgs *PGTaskStore) CreateTask(title string, description string, priority int32) (store.Task, error) {
	task, err := pgs.Queries.CreateTask(context.Background(), CreateTaskParams{
		Title:       title,
		Description: description,
		Priority:    priority,
	})
	if err != nil {
		return store.Task{}, err
	}
	return mapTask(task), nil
}

func (pgs *PGTaskStore) GetTaskById(id int32) (store.Task, error) {
	task, err := pgs.Queries.GetTaskById(context.Background(), id)
	if err != nil {
		return store.Task{}, err
	}
	return mapTask(task), nil
}

func (pgs *PGTaskStore) ListTasks() ([]store.Task, error) {
	tasks, err := pgs.Queries.ListTasks(context.Background())
	if err != nil {
		return nil, err
	}

	result := make([]store.Task, 0, len(tasks))
	for _, task := range tasks {
		result = append(result, mapTask(task))
	}
	return result, nil
}

func (pgs *PGTaskStore) UpdateTask(id int32, title string, description string, priority int32) (store.Task, error) {
	task, err := pgs.Queries.UpdateTask(context.Background(), UpdateTaskParams{
		Title:       title,
		Description: description,
		Priority:    priority,
		ID:          id,
	})
	if err != nil {
		return store.Task{}, err
	}
	return mapTask(task), nil
}

func (pgs *PGTaskStore) DeleteTask(id int32) error {
	return pgs.Queries.DeleteTask(context.Background(), id)
}
