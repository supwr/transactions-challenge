package transaction

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/supwr/pismo-transactions/internal/entity"
	accountservice "github.com/supwr/pismo-transactions/internal/usecase/account"
	accountrepo "github.com/supwr/pismo-transactions/internal/usecase/account/mock"
	operationtypeservice "github.com/supwr/pismo-transactions/internal/usecase/operation_type"
	operationtyperepo "github.com/supwr/pismo-transactions/internal/usecase/operation_type/mock"
	transactionrepo "github.com/supwr/pismo-transactions/internal/usecase/transaction/mock"
	clockmock "github.com/supwr/pismo-transactions/pkg/clock/mock"
	"testing"
	"time"
)

func TestService_Create(t *testing.T) {
	t.Run("create cash buy transaction successfully", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		accountRepo := accountrepo.NewMockRepositoryInterface(ctrl)
		operationTypeRepo := operationtyperepo.NewMockRepositoryInterface(ctrl)
		transactionRepo := transactionrepo.NewMockRepositoryInterface(ctrl)
		clockMock := clockmock.NewMockClock(ctrl)
		ctx := context.Background()

		account := &entity.Account{
			ID:       1,
			Document: "123456",
		}

		operationCashBuy := entity.OperationTypeCashBuy

		operationType := &entity.OperationType{
			ID:   operationCashBuy,
			Name: entity.Operations[operationCashBuy],
		}

		findAccountById := accountRepo.EXPECT().FindById(ctx, 1).Return(account, nil).Times(1)
		findOperationTypeById := operationTypeRepo.EXPECT().FindById(ctx, entity.OperationTypeCashBuy).Return(operationType, nil).Times(1).After(findAccountById)

		transactionDate := time.Now()
		transaction := &entity.Transaction{
			AccountID:       1,
			OperationTypeID: operationCashBuy,
			Amount:          decimal.NewFromFloat(float64(123.45)).Neg(),
			OperationDate:   transactionDate,
		}

		clock := clockMock.EXPECT().Now().Return(transactionDate).Times(1).After(findOperationTypeById)
		transactionRepo.EXPECT().Create(ctx, transaction).Return(nil).After(clock).Times(1)

		accountService := accountservice.NewService(accountRepo)
		operationTypeService := operationtypeservice.NewService(operationTypeRepo)
		transactionService := NewService(transactionRepo, operationTypeService, accountService, clockMock)

		err := transactionService.Create(ctx, &entity.Transaction{
			AccountID:       transaction.AccountID,
			OperationTypeID: transaction.OperationTypeID,
			Amount:          transaction.Amount.Abs(),
			OperationDate:   transactionDate,
		})

		assert.Nil(t, err)
	})

	t.Run("create payment transaction successfully", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		accountRepo := accountrepo.NewMockRepositoryInterface(ctrl)
		operationTypeRepo := operationtyperepo.NewMockRepositoryInterface(ctrl)
		transactionRepo := transactionrepo.NewMockRepositoryInterface(ctrl)
		clockMock := clockmock.NewMockClock(ctrl)
		ctx := context.Background()

		account := &entity.Account{
			ID:       1,
			Document: "123456",
		}

		operationPayment := entity.OperationTypePayment

		operationType := &entity.OperationType{
			ID:   operationPayment,
			Name: entity.Operations[operationPayment],
		}

		findAccountById := accountRepo.EXPECT().FindById(ctx, 1).Return(account, nil).Times(1)
		findOperationTypeById := operationTypeRepo.EXPECT().FindById(ctx, entity.OperationTypePayment).Return(operationType, nil).Times(1).After(findAccountById)

		transactionDate := time.Now()
		transaction := &entity.Transaction{
			AccountID:       1,
			OperationTypeID: operationPayment,
			Amount:          decimal.NewFromFloat(float64(123.45)),
			OperationDate:   transactionDate,
		}

		clock := clockMock.EXPECT().Now().Return(transactionDate).Times(1).After(findOperationTypeById)
		transactionRepo.EXPECT().Create(ctx, transaction).Return(nil).After(clock).Times(1)

		accountService := accountservice.NewService(accountRepo)
		operationTypeService := operationtypeservice.NewService(operationTypeRepo)
		transactionService := NewService(transactionRepo, operationTypeService, accountService, clockMock)

		err := transactionService.Create(ctx, &entity.Transaction{
			AccountID:       transaction.AccountID,
			OperationTypeID: transaction.OperationTypeID,
			Amount:          transaction.Amount.Abs(),
			OperationDate:   transactionDate,
		})

		assert.Nil(t, err)
	})

	t.Run("find account by id error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		accountRepo := accountrepo.NewMockRepositoryInterface(ctrl)
		operationTypeRepo := operationtyperepo.NewMockRepositoryInterface(ctrl)
		transactionRepo := transactionrepo.NewMockRepositoryInterface(ctrl)
		clockMock := clockmock.NewMockClock(ctrl)
		expectedError := errors.New("database error")
		transactionDate := time.Now()
		ctx := context.Background()

		transaction := &entity.Transaction{
			AccountID:       1,
			OperationTypeID: entity.OperationTypeCashBuy,
			Amount:          decimal.NewFromFloat(float64(123.45)).Neg(),
			OperationDate:   transactionDate,
		}

		accountRepo.EXPECT().FindById(ctx, 1).Return(nil, expectedError).Times(1)

		accountService := accountservice.NewService(accountRepo)
		operationTypeService := operationtypeservice.NewService(operationTypeRepo)
		transactionService := NewService(transactionRepo, operationTypeService, accountService, clockMock)

		err := transactionService.Create(ctx, &entity.Transaction{
			AccountID:       transaction.AccountID,
			OperationTypeID: transaction.OperationTypeID,
			Amount:          transaction.Amount.Abs(),
			OperationDate:   transactionDate,
		})

		assert.ErrorIs(t, err, expectedError)
	})

	t.Run("account not found error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		accountRepo := accountrepo.NewMockRepositoryInterface(ctrl)
		operationTypeRepo := operationtyperepo.NewMockRepositoryInterface(ctrl)
		transactionRepo := transactionrepo.NewMockRepositoryInterface(ctrl)
		clockMock := clockmock.NewMockClock(ctrl)
		transactionDate := time.Now()
		ctx := context.Background()

		transaction := &entity.Transaction{
			AccountID:       1,
			OperationTypeID: entity.OperationTypeCashBuy,
			Amount:          decimal.NewFromFloat(float64(123.45)).Neg(),
			OperationDate:   transactionDate,
		}

		accountRepo.EXPECT().FindById(ctx, 1).Return(nil, nil).Times(1)

		accountService := accountservice.NewService(accountRepo)
		operationTypeService := operationtypeservice.NewService(operationTypeRepo)
		transactionService := NewService(transactionRepo, operationTypeService, accountService, clockMock)

		err := transactionService.Create(ctx, &entity.Transaction{
			AccountID:       transaction.AccountID,
			OperationTypeID: transaction.OperationTypeID,
			Amount:          transaction.Amount.Abs(),
			OperationDate:   transactionDate,
		})

		assert.ErrorIs(t, err, ErrAccountNotFound)
	})

	t.Run("find operation type error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		accountRepo := accountrepo.NewMockRepositoryInterface(ctrl)
		operationTypeRepo := operationtyperepo.NewMockRepositoryInterface(ctrl)
		transactionRepo := transactionrepo.NewMockRepositoryInterface(ctrl)
		clockMock := clockmock.NewMockClock(ctrl)
		expectedError := errors.New("database error")
		ctx := context.Background()

		account := &entity.Account{
			ID:       1,
			Document: "123456",
		}

		findAccountById := accountRepo.EXPECT().FindById(ctx, 1).Return(account, nil).Times(1)
		operationTypeRepo.EXPECT().FindById(ctx, entity.OperationTypeCashBuy).Return(nil, expectedError).Times(1).After(findAccountById)

		transactionDate := time.Now()
		transaction := &entity.Transaction{
			AccountID:       1,
			OperationTypeID: entity.OperationTypeCashBuy,
			Amount:          decimal.NewFromFloat(float64(123.45)).Neg(),
			OperationDate:   transactionDate,
		}

		accountService := accountservice.NewService(accountRepo)
		operationTypeService := operationtypeservice.NewService(operationTypeRepo)
		transactionService := NewService(transactionRepo, operationTypeService, accountService, clockMock)

		err := transactionService.Create(ctx, &entity.Transaction{
			AccountID:       transaction.AccountID,
			OperationTypeID: transaction.OperationTypeID,
			Amount:          transaction.Amount.Abs(),
			OperationDate:   transactionDate,
		})

		assert.ErrorIs(t, err, expectedError)
	})

	t.Run("operation type not found error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		accountRepo := accountrepo.NewMockRepositoryInterface(ctrl)
		operationTypeRepo := operationtyperepo.NewMockRepositoryInterface(ctrl)
		transactionRepo := transactionrepo.NewMockRepositoryInterface(ctrl)
		clockMock := clockmock.NewMockClock(ctrl)
		ctx := context.Background()

		account := &entity.Account{
			ID:       1,
			Document: "123456",
		}

		findAccountById := accountRepo.EXPECT().FindById(ctx, 1).Return(account, nil).Times(1)
		operationTypeRepo.EXPECT().FindById(ctx, entity.OperationTypeCashBuy).Return(nil, nil).Times(1).After(findAccountById)

		transactionDate := time.Now()
		transaction := &entity.Transaction{
			AccountID:       1,
			OperationTypeID: entity.OperationTypeCashBuy,
			Amount:          decimal.NewFromFloat(float64(123.45)).Neg(),
			OperationDate:   transactionDate,
		}

		accountService := accountservice.NewService(accountRepo)
		operationTypeService := operationtypeservice.NewService(operationTypeRepo)
		transactionService := NewService(transactionRepo, operationTypeService, accountService, clockMock)

		err := transactionService.Create(ctx, &entity.Transaction{
			AccountID:       transaction.AccountID,
			OperationTypeID: transaction.OperationTypeID,
			Amount:          transaction.Amount.Abs(),
			OperationDate:   transactionDate,
		})

		assert.ErrorIs(t, err, ErrOperationTypeNotFound)
	})

	t.Run("create installment buy transaction successfully", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		accountRepo := accountrepo.NewMockRepositoryInterface(ctrl)
		operationTypeRepo := operationtyperepo.NewMockRepositoryInterface(ctrl)
		transactionRepo := transactionrepo.NewMockRepositoryInterface(ctrl)
		clockMock := clockmock.NewMockClock(ctrl)
		ctx := context.Background()

		account := &entity.Account{
			ID:       1,
			Document: "123456",
		}

		operationInstallmentBuy := entity.OperationTypeInstallmentBuy

		operationType := &entity.OperationType{
			ID:   operationInstallmentBuy,
			Name: entity.Operations[operationInstallmentBuy],
		}

		findAccountById := accountRepo.EXPECT().FindById(ctx, 1).Return(account, nil).Times(1)
		findOperationTypeById := operationTypeRepo.EXPECT().FindById(ctx, entity.OperationTypeInstallmentBuy).Return(operationType, nil).Times(1).After(findAccountById)

		transactionDate := time.Now()
		transaction := &entity.Transaction{
			AccountID:       1,
			OperationTypeID: operationInstallmentBuy,
			Amount:          decimal.NewFromFloat(float64(657.89)).Neg(),
			OperationDate:   transactionDate,
		}

		clock := clockMock.EXPECT().Now().Return(transactionDate).Times(1).After(findOperationTypeById)
		transactionRepo.EXPECT().Create(ctx, transaction).Return(nil).After(clock).Times(1)

		accountService := accountservice.NewService(accountRepo)
		operationTypeService := operationtypeservice.NewService(operationTypeRepo)
		transactionService := NewService(transactionRepo, operationTypeService, accountService, clockMock)

		err := transactionService.Create(ctx, &entity.Transaction{
			AccountID:       transaction.AccountID,
			OperationTypeID: transaction.OperationTypeID,
			Amount:          transaction.Amount.Abs(),
			OperationDate:   transactionDate,
		})

		assert.Nil(t, err)
	})

	t.Run("create withdraw transaction successfully", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		accountRepo := accountrepo.NewMockRepositoryInterface(ctrl)
		operationTypeRepo := operationtyperepo.NewMockRepositoryInterface(ctrl)
		transactionRepo := transactionrepo.NewMockRepositoryInterface(ctrl)
		clockMock := clockmock.NewMockClock(ctrl)
		ctx := context.Background()

		account := &entity.Account{
			ID:       1,
			Document: "123456",
		}

		operationWithdraw := entity.OperationTypeWithdraw

		operationType := &entity.OperationType{
			ID:   operationWithdraw,
			Name: entity.Operations[operationWithdraw],
		}

		findAccountById := accountRepo.EXPECT().FindById(ctx, 1).Return(account, nil).Times(1)
		findOperationTypeById := operationTypeRepo.EXPECT().FindById(ctx, entity.OperationTypeWithdraw).Return(operationType, nil).Times(1).After(findAccountById)

		transactionDate := time.Now()
		transaction := &entity.Transaction{
			AccountID:       1,
			OperationTypeID: operationWithdraw,
			Amount:          decimal.NewFromFloat(float64(654.32)).Neg(),
			OperationDate:   transactionDate,
		}

		clock := clockMock.EXPECT().Now().Return(transactionDate).Times(1).After(findOperationTypeById)
		transactionRepo.EXPECT().Create(ctx, transaction).Return(nil).After(clock).Times(1)

		accountService := accountservice.NewService(accountRepo)
		operationTypeService := operationtypeservice.NewService(operationTypeRepo)
		transactionService := NewService(transactionRepo, operationTypeService, accountService, clockMock)

		err := transactionService.Create(ctx, &entity.Transaction{
			AccountID:       transaction.AccountID,
			OperationTypeID: transaction.OperationTypeID,
			Amount:          transaction.Amount.Abs(),
			OperationDate:   transactionDate,
		})

		assert.Nil(t, err)
	})
}
