package dto

import "time"

type SalesOrderResponse struct {
	ID            uint       `json:"id"`
	PoBuyerNo     string     `json:"po_buyer_no"`
	OrderTypeID   uint       `json:"order_type_id"`
	CustomerID    uint       `json:"customer_id"`
	Status        string     `json:"status"`
	OrderAt       *time.Time `json:"order_at"`
	ShippingAt    *time.Time `json:"shipping_at"`
	CurrencyID    uint       `json:"currency_id"`
	ExchangeRate  float64    `json:"exchange_rate"`
	IsVat         bool       `json:"is_vat"`
	IsPph23       bool       `json:"is_pph23"`
	VatID         uint       `json:"vat_id"`
	Pph23ID       uint       `json:"pph23_id"`
	Subtotal      float64    `json:"sub_total"`
	TotalDiscount float64    `json:"total_discount"`
	TotalVat      float64    `json:"total_vat"`
	TotalPph23    float64    `json:"total_pph23"`
	GrandTotal    float64    `json:"grand_total"`
	Remark        string     `json:"remark"`
}
