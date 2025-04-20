package repositories

import (
	"errors"
	"yubi-fullstack-test/models"

	"gorm.io/gorm"
)

type SoDtRepository interface {
	FindAllBySalesOrderId(salesOrderId uint) ([]*models.SoDt, string)
	FindAllBySalesOrderIdAndDetailId(salesOrderId uint, detailId uint) (*models.SoDt, string)
	Store(soDt models.SoDt) (*models.SoDt, string)
	Update(soDt models.SoDt) (*models.SoDt, string)
	Delete(id uint) (string, string)
}

type soDtRepository struct {
	db *gorm.DB
}

func NewSoDtRepository(db *gorm.DB) SoDtRepository {
	return &soDtRepository{db: db}
}

func (r *soDtRepository) FindAllBySalesOrderId(salesOrderId uint) ([]*models.SoDt, string) {
	var soDts []*models.SoDt

	if err := r.db.Where("sales_order_id = ?", salesOrderId).Find(&soDts).Error; err != nil {
		return nil, "DATABASE_ERROR_500"
	}

	return soDts, ""
}

func (r *soDtRepository) FindAllBySalesOrderIdAndDetailId(salesOrderId uint, detailId uint) (*models.SoDt, string) {
	var soDt *models.SoDt

	if err := r.db.Where("sales_order_id = ? AND id = ?", salesOrderId, detailId).First(&soDt).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, "SO_DT_NOT_FOUND_404"
		}
		return nil, "DATABASE_ERROR_500"
	}

	return soDt, ""
}

func (r *soDtRepository) Store(soDt models.SoDt) (*models.SoDt, string) {
	if err := r.db.Create(&soDt).Error; err != nil {
		return nil, "DATABASE_ERROR_500"
	}

	return &soDt, ""

}

func (r *soDtRepository) Update(soDt models.SoDt) (*models.SoDt, string) {
	if err := r.db.Save(&soDt).Error; err != nil {
		return nil, "DATABASE_ERROR_500"
	}

	return &soDt, ""

}

func (r *soDtRepository) Delete(id uint) (string, string) {
	var soDt *models.SoDt

	if err := r.db.First(&soDt, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", "SO_DT_NOT_FOUND_404"
		}
		return "", "DATABASE_ERROR_500"
	}

	if err := r.db.Delete(&soDt).Error; err != nil {
		return "", "DATABASE_ERROR_500"
	}

	return "DELETED", ""

}
