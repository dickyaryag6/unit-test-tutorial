package usecase

import (
	"errors"
	"gomocking/repository"
	"strconv"
)

type Usecase struct {
	Database repository.DatabaseInterface
}

func New(usecase Usecase) Usecase {
	return usecase
}

func (u *Usecase) Addition(number1, number2 string) (int, error) {
	
	number1Int, err := convertStringToInt(number1)
	if err != nil {
		return 0, err
	}

	number2Int, err := convertStringToInt(number2)
	if err != nil {
		return 0, err
	}

	result := number1Int + number2Int

	err = u.Database.Insert(result)
	if err != nil {
		return 0, err
	}

	return result, nil
}

func convertStringToInt(s string) (int, error) {
	n, err := strconv.Atoi(s)
	
	if err != nil {
		return 0, errors.New("Input can't be converted to integer")
	}
		
	return n, nil
	
}