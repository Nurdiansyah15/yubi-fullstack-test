package models

import (
	"time"
)

type SalesOrder struct {
	ID            uint       `gorm:"primaryKey" json:"id"`
	CustomerID    uint       `json:"customer_id"`
	OrderTypeID   uint       `json:"order_type_id"`
	CurrencyID    uint       `json:"currency_id"`
	WarehouseID   uint       `json:"warehouse_id"`
	PaymentID     uint       `json:"payment_id"`
	VatID         uint       `json:"vat_id"`
	Pph23ID       uint       `json:"pph23_id"`
	RevNo         string     `json:"rev_no"`
	NoBuyerNo     string     `json:"no_buyer_no"`
	NoBuyerNoOri  string     `json:"no_buyer_no_ori"`
	SalesOrderNo  string     `json:"sales_order_no"`
	Remark        string     `json:"remark"`
	ShipDest      string     `json:"ship_dest"`
	Status        string     `json:"status"`
	ExchangeRate  float64    `json:"exchange_rate"`
	VatPerc       float64    `json:"vat_perc"`
	Pph23Perc     float64    `json:"pph23_perc"`
	MarkupPerc    float64    `json:"markup_perc"`
	IsVat         bool       `json:"is_vat"`
	IsPph23       bool       `json:"is_pph23"`
	DiscAm        float64    `json:"disc_am"`
	DiscPerc      float64    `json:"disc_perc"`
	DiscFinal     float64    `json:"disc_final"`
	TotalQty      float64    `json:"total_qty"`
	Subtotal      float64    `json:"subtotal"`
	TotalDiscount float64    `json:"total_discount"`
	TotalPph23    float64    `json:"total_pph23"`
	TotalVat      float64    `json:"total_vat"`
	GrandTotal    float64    `json:"grand_total"`
	QtyOut        float64    `json:"qty_out"`
	SaTotalAm     float64    `json:"sa_total_am"`
	Total         float64    `json:"total"`
	OrderAt       *time.Time `json:"order_at"`
	ShippingAt    *time.Time `json:"shipping_at"`
	DueAt         *time.Time `json:"due_at"`
	ExpiredAt     *time.Time `json:"expired_at"`
}
