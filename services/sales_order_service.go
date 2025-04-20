package services

import (
	"yubi-fullstack-test/dto"
	"yubi-fullstack-test/models"
	"yubi-fullstack-test/repositories"
	formatterUtils "yubi-fullstack-test/utils/formatter"
)

type SalesOrderService interface {
	FindAll() ([]*dto.SalesOrderResponse, string)
	FindById(id uint) (*dto.SalesOrderResponse, string)
	Store(salesOrder models.SalesOrder) (*models.SalesOrder, string)
	Update(salesOrder models.SalesOrder) (*models.SalesOrder, string)
	Delete(id uint) (string, string)
}

type salesOrderService struct {
	salesOrderRepository repositories.SalesOrderRepository
}

func NewSalesOrderService(salesOrderRepository repositories.SalesOrderRepository) SalesOrderService {
	return &salesOrderService{salesOrderRepository: salesOrderRepository}
}

func (s *salesOrderService) FindAll() ([]*dto.SalesOrderResponse, string) {
	salesOrders, err := s.salesOrderRepository.FindAll()
	if err != "" {
		return nil, err
	}

	salesOrdersResponse := formatterUtils.FormatSalesOrders(salesOrders)

	return salesOrdersResponse, ""
}

func (s *salesOrderService) FindById(id uint) (*dto.SalesOrderResponse, string) {
	salesOrder, err := s.salesOrderRepository.FindById(id)
	if err != "" {
		return nil, err
	}

	salesOrderResponse := formatterUtils.FormatSalesOrder(salesOrder)

	return salesOrderResponse, ""
}

func (s *salesOrderService) Store(salesOrder models.SalesOrder) (*models.SalesOrder, string) {
	salesOrderStored, err := s.salesOrderRepository.Store(salesOrder)
	if err != "" {
		return nil, err
	}

	return salesOrderStored, ""
}

func (s *salesOrderService) Update(salesOrder models.SalesOrder) (*models.SalesOrder, string) {
	salesOrderUpdated, err := s.salesOrderRepository.Update(salesOrder)
	if err != "" {
		return nil, err
	}

	return salesOrderUpdated, ""
}

func (s *salesOrderService) Delete(id uint) (string, string) {
	result, err := s.salesOrderRepository.Delete(id)
	if err != "" {
		return "", err
	}

	return result, ""
}
