package memdb

import "30-8-1/pkg/storage/postgresql"

type DB []postgresql.Task

func (db DB) Tasks(int, int) ([]postgresql.Task, error) {
	return db, nil
}
func (db DB) NewTask(postgresql.Task) (int, error) {
	return 0, nil
}

func (db DB) DelTask(postgresql.Task) error {
	return nil
}

func (db DB) UpdateTask(postgresql.Task) error {
	return nil
}

func (db DB) TasksbyLables(string) ([]postgresql.Task, error) {
	return db, nil
}
