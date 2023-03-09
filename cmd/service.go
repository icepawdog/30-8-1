package main

import (
	"30-8-1/pkg/storage"
	"30-8-1/pkg/storage/postgresql"
	"fmt"
	"log"
)

var db storage.Interface

func main() {
	var err error
	connstr := "postgres://postgres:password@server.domain/tasks"
	db, err = postgresql.New(connstr)
	if err != nil {
		log.Fatal(err)
	}
	// db = memdb.DB{}
	tasks, err := db.TasksbyLables("крутая задача")
	if err != nil {
		log.Fatal(err)
	}
	tasks2, err := db.NewTask(postgresql.Task{1, 2, 3, 4, 5, "sa", "oa"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tasks2)
	fmt.Println(tasks)
}
