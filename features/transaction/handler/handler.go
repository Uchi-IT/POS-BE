package handler

import (
	"net/http"
	"strconv"
	"uchiiParfume/features/transaction/dto"
	"uchiiParfume/features/transaction/entity"
	middleware "uchiiParfume/utils/jwt"

	"github.com/labstack/echo/v4"
)

type transactionHandler struct {
	transactionService entity.TransactionServiceInterface
}

func NewTransactionHandler(transaction entity.TransactionServiceInterface) *transactionHandler {
	return &transactionHandler{
		transactionService: transaction,
	}
}

func (transaction *transactionHandler) CreateTransaction(e echo.Context) error {
	userId, _, err := middleware.ExtractToken(e)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": err.Error(),
		})
	}

	input := dto.TransactionRequest{}
	errBind := e.Bind(&input)
	if errBind != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind data",
		})
	}

	transactionInput := dto.TransactionRequestToTransactionCore(input)

	_, errUser := transaction.transactionService.CreateTransaction(transactionInput, userId)
	if errUser != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error create transaction",
			"error":   errUser.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]any{
		"message": "succes create transaction",
	})
}

func (transaction *transactionHandler) GetAllTransaction(e echo.Context) error {
	data, err := transaction.transactionService.GetAllTransaction()
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error get all transaction",
		})
	}

	dataList := dto.ListTransactionCoreToListTransactionResponse(data)

	return e.JSON(http.StatusOK, map[string]any{
		"message": "get all transaction",
		"data":    dataList,
	})
}

func (transaction *transactionHandler) GetAllTransactionItemById(e echo.Context) error {
	idParamstr := e.Param("id")

	idInt, errConv := strconv.Atoi(idParamstr)
	if errConv != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "invalid transaction id",
		})
	}

	data, err := transaction.transactionService.GetAllTransactionItemById(idInt)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error get all transaction item",
		})
	}

	dataList := dto.ListTransactionItemCoreToListTransactionItemResponse(data)

	return e.JSON(http.StatusOK, map[string]any{
		"message": "get all	 transaction Items",
		"data":    dataList,
	})
}

func (transaction *transactionHandler) GetSpecificTransaction(e echo.Context) error {
	idParamstr := e.Param("id")

	idInt, errConv := strconv.Atoi(idParamstr)
	if errConv != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "invalid transaction id",
		})
	}

	data, err := transaction.transactionService.GetSpecificTransaction(idInt)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error get specific transaction",
		})
	}

	dataSpecific := dto.TransactionCoreToTransactionResponse(data)

	return e.JSON(http.StatusOK, map[string]any{
		"message": "get specific transaction",
		"data":    dataSpecific,
	})
}
