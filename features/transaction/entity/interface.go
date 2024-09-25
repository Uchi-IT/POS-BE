package entity

type TransactionRepositoryInterface interface {
	CreateTransaction(data TransactionCore, userId string) (TransactionCore, error)
	GetAllTransaction() ([]TransactionCore, error)
	GetSpecificTransaction(id int) (TransactionCore, error)
	GetAllTransactionItemById(id int) ([]TransactionItemCore, error)
}

type TransactionServiceInterface interface{
	CreateTransaction(data TransactionCore, userId string) (TransactionCore, error)
	GetAllTransaction() ([]TransactionCore, error)
	GetSpecificTransaction(id int) (TransactionCore, error)
	GetAllTransactionItemById(id int) ([]TransactionItemCore, error)
}