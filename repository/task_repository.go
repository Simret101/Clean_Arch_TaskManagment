package repository

import (
	"context"
	"errors"
	"strings"
	"sync"
	"task/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InMemoryTaskRepository struct {
	tasks  []domain.Task
	taskMu sync.Mutex
}

func NewInMemoryTaskRepository() *InMemoryTaskRepository {
	return &InMemoryTaskRepository{tasks: []domain.Task{}}
}

func (r *InMemoryTaskRepository) CreateTask(ctx context.Context, task *domain.Task) error {
	r.taskMu.Lock()
	defer r.taskMu.Unlock()

	if err := r.Validate(task); err != nil {
		return err
	}

	task.ID = primitive.NewObjectID()

	r.tasks = append(r.tasks, *task)
	return nil
}

func (r *InMemoryTaskRepository) GetAllTasks(ctx context.Context) ([]domain.Task, error) {
	r.taskMu.Lock()
	defer r.taskMu.Unlock()

	return r.tasks, nil
}

func (r *InMemoryTaskRepository) GetTaskByID(ctx context.Context, id primitive.ObjectID) (*domain.Task, error) {
	r.taskMu.Lock()
	defer r.taskMu.Unlock()

	for _, task := range r.tasks {
		if task.ID == id {
			return &task, nil
		}
	}
	return nil, errors.New("task not found")
}

func (r *InMemoryTaskRepository) UpdateTask(ctx context.Context, task *domain.Task) error {
	r.taskMu.Lock()
	defer r.taskMu.Unlock()

	if err := r.Validate(task); err != nil {
		return err
	}

	for i, t := range r.tasks {
		if t.ID == task.ID {
			r.tasks[i] = *task
			return nil
		}
	}
	return errors.New("task not found")
}

func (r *InMemoryTaskRepository) DeleteTask(ctx context.Context, id primitive.ObjectID) error {
	r.taskMu.Lock()
	defer r.taskMu.Unlock()

	for i, task := range r.tasks {
		if task.ID == id {
			r.tasks = append(r.tasks[:i], r.tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}

func (r *InMemoryTaskRepository) Validate(t *domain.Task) error {
	if err := r.validateTitle(t.Title); err != nil {
		return err
	}
	if err := r.validateDescription(t.Description); err != nil {
		return err
	}
	if err := r.validateStatus(t.Status); err != nil {
		return err
	}
	return nil
}

func (r *InMemoryTaskRepository) validateTitle(title string) error {
	if strings.TrimSpace(title) == "" {
		return errors.New("title must not be empty")
	}
	if len(title) > 100 {
		return errors.New("title must be less than 100 characters")
	}
	if len(title) < 3 {
		return errors.New("title must be greater than 3 characters")
	}
	return nil
}

func (r *InMemoryTaskRepository) validateDescription(description string) error {
	if strings.TrimSpace(description) == "" {
		return errors.New("description must not be empty")
	}
	return nil
}

func (r *InMemoryTaskRepository) validateStatus(status domain.TaskStatus) error {
	switch status {
	case domain.TaskStatusComplete, domain.TaskStatusInProgress, domain.TaskStatusStarted:
		return nil
	default:
		return errors.New("status is invalid")
	}
}
