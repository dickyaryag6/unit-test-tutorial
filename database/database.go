package repository

import (
	"fmt"
)

type Database struct {}

func New(db Database) Database {
	return db
}

func (db Database) Insert(value int) error {
	fmt.Println("Inserting to database ...")

	return nil
}