package services

import (
	"yubi-fullstack-test/dto"
	"yubi-fullstack-test/models"
	"yubi-fullstack-test/repositories"
	formatterUtils "yubi-fullstack-test/utils/formatter"
)

type SoDtService interface {
	FindAllBySalesOrderId(salesOrderId uint) ([]*dto.SalesOrderDetailResponse, string)
	FindAllBySalesOrderIdAndDetailId(salesOrderId uint, detailId uint) (*dto.SalesOrderDetailResponse, string)
	Store(salesOrderDetail models.SoDt) (*models.SoDt, string)
	Update(salesOrderDetail models.SoDt) (*models.SoDt, string)
	Delete(id uint) (string, string)
}

type soDtService struct {
	soDtRepository repositories.SoDtRepository
	soService      SalesOrderService
}

func NewSoDtService(soDtRepository repositories.SoDtRepository, soService SalesOrderService) SoDtService {
	return &soDtService{soDtRepository: soDtRepository, soService: soService}
}

func (s *soDtService) FindAllBySalesOrderId(salesOrderId uint) ([]*dto.SalesOrderDetailResponse, string) {
	soDts, err := s.soDtRepository.FindAllBySalesOrderId(salesOrderId)
	if err != "" {
		return nil, err
	}

	soDtsResponse := formatterUtils.FormatSalesOrderDetails(soDts)

	return soDtsResponse, ""
}

func (s *soDtService) FindAllBySalesOrderIdAndDetailId(salesOrderId uint, detailId uint) (*dto.SalesOrderDetailResponse, string) {
	soDt, err := s.soDtRepository.FindAllBySalesOrderIdAndDetailId(salesOrderId, detailId)

	if err != "" {
		return nil, err
	}

	soDtResponse := formatterUtils.FormatSalesOrderDetail(soDt)

	return soDtResponse, ""
}

func (s *soDtService) Store(soDt models.SoDt) (*models.SoDt, string) {

	_, err := s.soService.FindById(soDt.SalesOrderID)
	if err != "" {
		return nil, err
	}

	soDtStored, err := s.soDtRepository.Store(soDt)
	if err != "" {
		return nil, err
	}

	return soDtStored, ""
}

func (s *soDtService) Update(soDt models.SoDt) (*models.SoDt, string) {

	_, err := s.soService.FindById(soDt.SalesOrderID)
	if err != "" {
		return nil, err
	}

	soDtUpdated, err := s.soDtRepository.Update(soDt)
	if err != "" {
		return nil, err
	}

	return soDtUpdated, ""
}

func (s *soDtService) Delete(id uint) (string, string) {
	_, err := s.soDtRepository.Delete(id)
	if err != "" {
		return "", err
	}

	return "DELETED", ""
}
