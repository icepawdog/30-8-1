package storage

import "30-8-1/pkg/storage/postgresql"

type Interface interface {
	Tasks(int, int) ([]postgresql.Task, error)
	NewTask(postgresql.Task) (int, error)
	UpdateTask(postgresql.Task) error
	DelTask(postgresql.Task) error
	TasksbyLables(string) ([]postgresql.Task, error)
}
