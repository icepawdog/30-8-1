package main

import (
	"30-8-1/pkg/storage"
	"30-8-1/pkg/storage/memdb"
	"30-8-1/pkg/storage/postgresql"
	"fmt"
	"log"
)

var db storage.Interface

func main() {
	var err error
	connstr := "postgres://postgres:password@server.domain/posts"
	db, err = postgresql.New(connstr)
	if err != nil {
		log.Fatal(err)
	}
	db = memdb.DB{}
	tasks, err := db.Tasks(0, 0)
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
