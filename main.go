package main

import (
	"fmt"
	db "gomocking/database"
	uc "gomocking/usecase"
)

func main() {

	databaseInstance := db.New(db.Database{})

	usecaseInstance := uc.New(uc.Usecase{
		Database: databaseInstance,
	})

	result, err := usecaseInstance.Addition("60", "20")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

}
