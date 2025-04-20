package repositories

import (
	"errors"
	"yubi-fullstack-test/models"

	"gorm.io/gorm"
)

type SalesOrderRepository interface {
	FindAll() ([]*models.SalesOrder, string)
	FindById(id uint) (*models.SalesOrder, string)
	Store(salesOrder models.SalesOrder) (*models.SalesOrder, string)
	Update(salesOrder models.SalesOrder) (*models.SalesOrder, string)
	Delete(id uint) (string, string)
}

type salesOrderRepository struct {
	db *gorm.DB
}

func NewSalesOrderRepository(db *gorm.DB) SalesOrderRepository {
	return &salesOrderRepository{db: db}
}

func (r *salesOrderRepository) FindAll() ([]*models.SalesOrder, string) {
	var salesOrders []*models.SalesOrder

	if err := r.db.Find(&salesOrders).Error; err != nil {
		return nil, "DATABASE_ERROR_500"
	}

	

	return salesOrders, ""
}

func (r *salesOrderRepository) FindById(id uint) (*models.SalesOrder, string) {
	var salesOrder *models.SalesOrder

	if err := r.db.First(&salesOrder, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, "SALES_ORDER_NOT_FOUND_404"
		}
		return nil, "DATABASE_ERROR_500"
	}

	return salesOrder, ""
}

func (r *salesOrderRepository) Store(salesOrder models.SalesOrder) (*models.SalesOrder, string) {
	if err := r.db.Create(&salesOrder).Error; err != nil {
		return nil, "DATABASE_ERROR_500"
	}

	return &salesOrder, ""
}

func (r *salesOrderRepository) Update(salesOrder models.SalesOrder) (*models.SalesOrder, string) {
	if err := r.db.Save(&salesOrder).Error; err != nil {
		return nil, "DATABASE_ERROR_500"
	}

	return &salesOrder, ""
}

func (r *salesOrderRepository) Delete(id uint) (string, string) {
	var salesOrder *models.SalesOrder

	if err := r.db.First(&salesOrder, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", "SALES_ORDER_NOT_FOUND_404"
		}
		return "", "DATABASE_ERROR_500"
	}

	if err := r.db.Delete(&salesOrder).Error; err != nil {
		return "", "DATABASE_ERROR_500"
	}

	return "DELETED", ""
}
