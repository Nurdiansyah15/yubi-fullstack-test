package formatterUtils

import (
	"yubi-fullstack-test/dto"
	"yubi-fullstack-test/models"
)

func FormatSalesOrder(order *models.SalesOrder) *dto.SalesOrderResponse {
	return &dto.SalesOrderResponse{
		ID:            order.ID,
		PoBuyerNo:     order.PoBuyerNo,
		OrderTypeID:   order.OrderTypeID,
		CustomerID:    order.CustomerID,
		Status:        order.Status,
		OrderAt:       order.OrderAt,
		ShippingAt:    order.ShippingAt,
		CurrencyID:    order.CurrencyID,
		ExchangeRate:  order.ExchangeRate,
		IsVat:         order.IsVat,
		IsPph23:       order.IsPph23,
		VatID:         order.VatID,
		Pph23ID:       order.Pph23ID,
		Subtotal:      order.Subtotal,
		TotalDiscount: order.TotalDiscount,
		TotalVat:      order.TotalVat,
		TotalPph23:    order.TotalPph23,
		GrandTotal:    order.GrandTotal,
		Remark:        order.Remark,
	}
}

func FormatSalesOrders(orders []*models.SalesOrder) []*dto.SalesOrderResponse {
	var result []*dto.SalesOrderResponse
	for _, order := range orders {
		result = append(result, FormatSalesOrder(order))
	}
	return result
}
