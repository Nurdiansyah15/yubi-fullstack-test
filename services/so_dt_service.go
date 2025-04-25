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
	Store(salesOrderDetail models.SoDt) (*dto.SalesOrderDetailResponse, string)
	Update(salesOrderDetail models.SoDt) (*dto.SalesOrderDetailResponse, string)
	Delete(salesOrderId uint, soDtId uint) (string, string)
}

type soDtService struct {
	soDtRepository repositories.SoDtRepository
	soRepository   repositories.SalesOrderRepository
	soService      SalesOrderService
}

func NewSoDtService(soDtRepository repositories.SoDtRepository, soService SalesOrderService, soRepository repositories.SalesOrderRepository) SoDtService {
	return &soDtService{soDtRepository: soDtRepository, soService: soService, soRepository: soRepository}
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

func (s *soDtService) Store(soDt models.SoDt) (*dto.SalesOrderDetailResponse, string) {

	_, err := s.soService.FindById(soDt.SalesOrderID)
	if err != "" {
		return nil, err
	}

	subTotal := soDt.Qty * soDt.PriceSell

	var discAm float64
	if soDt.DiscPerc > 0 {
		discAm = subTotal * (soDt.DiscPerc / 100)
	} else {
		discAm = soDt.DiscAm
	}

	soDt.Total = subTotal - discAm

	soDtStored, err := s.soDtRepository.Store(soDt)
	if err != "" {
		return nil, err
	}

	if err := s.recalculateSalesOrderTotals(soDt.SalesOrderID); err != "" {
		return nil, err
	}

	soDtResponse := formatterUtils.FormatSalesOrderDetail(soDtStored)

	return soDtResponse, ""
}

func (s *soDtService) Update(soDt models.SoDt) (*dto.SalesOrderDetailResponse, string) {

	_, err := s.soService.FindById(soDt.SalesOrderID)
	if err != "" {
		return nil, err
	}

	subTotal := soDt.Qty * soDt.PriceSell

	var discAm float64
	if soDt.DiscPerc > 0 {
		discAm = subTotal * (soDt.DiscPerc / 100)
	} else {
		discAm = soDt.DiscAm
	}

	soDt.Total = subTotal - discAm

	soDtUpdated, err := s.soDtRepository.Update(soDt)
	if err != "" {
		return nil, err
	}

	if err := s.recalculateSalesOrderTotals(soDt.SalesOrderID); err != "" {
		return nil, err
	}

	soDtResponse := formatterUtils.FormatSalesOrderDetail(soDtUpdated)

	return soDtResponse, ""
}

func (s *soDtService) Delete(salesOrderId, soDtId uint) (string, string) {
	soDt, errStr := s.soDtRepository.FindAllBySalesOrderIdAndDetailId(salesOrderId, soDtId)
	if errStr != "" {
		return "", errStr
	}

	_, delErrStr := s.soDtRepository.Delete(soDt.ID)
	if delErrStr != "" {
		return "", delErrStr
	}

	recalcErrStr := s.recalculateSalesOrderTotals(soDt.SalesOrderID)
	if recalcErrStr != "" {
		return "", recalcErrStr
	}

	return "DELETED", ""
}

func (s *soDtService) recalculateSalesOrderTotals(salesOrderID uint) string {
	soDts, err := s.soDtRepository.FindAllBySalesOrderId(salesOrderID)
	if err != "" {
		return err
	}

	so, err := s.soRepository.FindById(salesOrderID)
	if err != "" {
		return err
	}

	var (
		subtotal      float64
		totalDiscount float64
	)

	for _, detail := range soDts {
		itemTotal := detail.PriceSell * detail.Qty

		var disc float64
		if detail.DiscAm > 0 {
			disc = detail.DiscAm
		} else if detail.DiscPerc > 0 {
			disc = itemTotal * (detail.DiscPerc / 100)
		}

		subtotal += itemTotal
		totalDiscount += disc
	}

	grandTotal := subtotal - totalDiscount

	so.Subtotal = subtotal
	so.TotalDiscount = totalDiscount
	so.GrandTotal = grandTotal

	_, err = s.soService.Update(*so)
	if err != "" {
		return err
	}

	return ""
}
