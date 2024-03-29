package repository

import (
	"context"
	"errors"
	"github.com/supwr/pismo-transactions/internal/entity"
	"gorm.io/gorm"
	"log/slog"
)

type OperationTypeRepository struct {
	db     *gorm.DB
	logger *slog.Logger
}

func NewOperationTypeRepository(db *gorm.DB, logger *slog.Logger) *OperationTypeRepository {
	return &OperationTypeRepository{
		db:     db,
		logger: logger,
	}
}

func (o *OperationTypeRepository) FindById(ctx context.Context, id int) (*entity.OperationType, error) {
	var operationType *entity.OperationType

	if err := o.db.First(&operationType, "id = ? and deleted_at is null", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		o.logger.ErrorContext(ctx, "error finding transaction type", slog.Any("error", err))
		return nil, err
	}

	return operationType, nil
}
