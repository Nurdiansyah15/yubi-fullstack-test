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
	Store(salesOrder models.SalesOrder) (*dto.SalesOrderResponse, string)
	Update(salesOrder models.SalesOrder) (*dto.SalesOrderResponse, string)
	Delete(id uint) (string, string)
	SetSoDtService(soDtService SoDtService)
}

type salesOrderService struct {
	salesOrderRepository repositories.SalesOrderRepository
	soDtsRepository      repositories.SoDtRepository
	soDtService          SoDtService
}

func NewSalesOrderService(salesOrderRepository repositories.SalesOrderRepository, soDtsRepository repositories.SoDtRepository) SalesOrderService {
	return &salesOrderService{salesOrderRepository: salesOrderRepository, soDtsRepository: soDtsRepository}
}

func (salesOrderService *salesOrderService) SetSoDtService(soDtService SoDtService) {
	salesOrderService.soDtService = soDtService
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

func (s *salesOrderService) Store(salesOrder models.SalesOrder) (*dto.SalesOrderResponse, string) {
	salesOrderStored, err := s.salesOrderRepository.Store(salesOrder)
	if err != "" {
		return nil, err
	}

	salesOrderResponse := formatterUtils.FormatSalesOrder(salesOrderStored)

	return salesOrderResponse, ""
}

func (s *salesOrderService) Update(salesOrder models.SalesOrder) (*dto.SalesOrderResponse, string) {
	salesOrderUpdated, err := s.salesOrderRepository.Update(salesOrder)
	if err != "" {
		return nil, err
	}
	salesOrderResponse := formatterUtils.FormatSalesOrder(salesOrderUpdated)

	return salesOrderResponse, ""
}

func (s *salesOrderService) Delete(id uint) (string, string) {
	soDts, err := s.soDtsRepository.FindAllBySalesOrderId(id)
	if err != "" {
		return "", err
	}

	for _, soDt := range soDts {
		_, err := s.soDtService.Delete(soDt.SalesOrderID, soDt.ID)
		if err != "" {
			return "", err
		}
	}

	result, err := s.salesOrderRepository.Delete(id)
	if err != "" {
		return "", err
	}

	return result, ""
}
