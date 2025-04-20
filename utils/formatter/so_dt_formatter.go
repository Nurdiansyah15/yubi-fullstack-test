package formatterUtils

import (
	"yubi-fullstack-test/dto"
	"yubi-fullstack-test/models"
)

func FormatSalesOrderDetail(detail *models.SoDt) *dto.SalesOrderDetailResponse {
	return &dto.SalesOrderDetailResponse{
		ID:           detail.ID,
		SalesOrderID: detail.SalesOrderID,
		RefType:      detail.RefType,
		ItemType:     detail.ItemType,
		ProductUUID:  detail.ProductUUID,
		ItemUnitID:   detail.ItemUnitID,
		PriceSell:    detail.PriceSell,
		Qty:          detail.Qty,
		DiscPerc:     detail.DiscPerc,
		DiscAm:       detail.DiscAm,
		TotalAm:      detail.Total,
		Remark:       detail.Remark,
	}
}

func FormatSalesOrderDetails(details []*models.SoDt) []*dto.SalesOrderDetailResponse {
	var result []*dto.SalesOrderDetailResponse

	for _, detail := range details {
		result = append(result, FormatSalesOrderDetail(detail))
	}

	return result
}
