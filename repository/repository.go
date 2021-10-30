package repository

//go:generate mockgen -destination=../mocks/mock_db.go -package=mocks gomocking/repository DatabaseInterface 

type DatabaseInterface interface {
	Insert(number int) error
}