package service

import (
	"errors"
	"uchiiParfume/features/transaction/entity"
)

type transactionService struct {
	TransactionRepository entity.TransactionRepositoryInterface
}

func NewTransactionService(transaction entity.TransactionRepositoryInterface) entity.TransactionServiceInterface {
	return &transactionService{
		TransactionRepository: transaction,
	}
}

// CreateTransaction implements entity.TransactionServiceInterface.
func (transUC *transactionService) CreateTransaction(data entity.TransactionCore, userId string) (entity.TransactionCore, error) {
	createData, err := transUC.TransactionRepository.CreateTransaction(data, userId)
	if err != nil {
		return entity.TransactionCore{}, err
	}

	return createData, nil
}

// GetAllTransaction implements entity.TransactionServiceInterface.
func (transUC *transactionService) GetAllTransaction() ([]entity.TransactionCore, error) {
	transactions, err := transUC.TransactionRepository.GetAllTransaction()
	if err != nil {
		return nil, errors.New("error get data")
	}

	return transactions, nil
}

// GetAllTransactionItemById implements entity.TransactionServiceInterface.
func (transUC *transactionService) GetAllTransactionItemById(id int) ([]entity.TransactionItemCore, error) {
	if id == 0 {
		return []entity.TransactionItemCore{}, errors.New("Transaction ID is required")
	}

	transactionItem, err := transUC.TransactionRepository.GetAllTransactionItemById(id)
	if err != nil {
		return []entity.TransactionItemCore{}, err
	}

	return transactionItem, nil
}

// GetSpecificTransaction implements entity.TransactionServiceInterface.
func (transUC *transactionService) GetSpecificTransaction(id int) (entity.TransactionCore, error) {
	if id == 0 {
		return entity.TransactionCore{}, errors.New("Transaction ID is required")
	}

	transaction, err := transUC.TransactionRepository.GetSpecificTransaction(id)
	if err != nil {
		return entity.TransactionCore{}, err
	}

	return transaction, nil
}