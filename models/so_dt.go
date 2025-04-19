package models

type SoDt struct {
	ID              uint    `gorm:"primaryKey" json:"id"`
	ProductUUID     string  `json:"product_uuid"`
	SalesOrderID    uint    `json:"sales_order_id"`
	ItemUnitID      uint    `json:"item_unit_id"`
	VatID           uint    `json:"vat_id"`
	Pph23ID         uint    `json:"pph23_id"`
	RefID           uint    `json:"ref_id"`
	RefJSON         string  `json:"ref_json"`
	RefType         string  `json:"ref_type"`
	ItemType        string  `json:"item_type"`
	ItemJSON        string  `json:"item_json"`
	GenCode         string  `json:"gen_code"`
	Remark          string  `json:"remark"`
	VatPercAm       float64 `json:"vat_perc_am"`
	Pph23PercAm     float64 `json:"pph23_perc_am"`
	MarkupPercAm    float64 `json:"markup_perc_am"`
	VatPerc         float64 `json:"vat_perc"`
	Pph23Perc       float64 `json:"pph23_perc"`
	MarkupPerc      float64 `json:"markup_perc"`
	IsVat           bool    `json:"is_vat"`
	IsPph23         bool    `json:"is_pph23"`
	IsLockMarkup    bool    `json:"is_lock_markup"`
	IsLockPriceSell bool    `json:"is_lock_price_sell"`
	QtyOut          float64 `json:"qty_out"`
	Qty             float64 `json:"qty"`
	PriceSell       float64 `json:"price_sell"`
	PriceBuy        float64 `json:"price_buy"`
	SubtotalSell    float64 `json:"subtotal_sell"`
	SubtotalBuy     float64 `json:"subtotal_buy"`
	DiscAm          float64 `json:"disc_am"`
	DiscPerc        float64 `json:"disc_perc"`
	DiscPercNum     float64 `json:"disc_perc_num"`
	DiscFinal       float64 `json:"disc_final"`
	Subtotal        float64 `json:"subtotal"`
	SaTotalAm       float64 `json:"sa_total_am"`
	Total           float64 `json:"total"`
	TotalDp         float64 `json:"total_dp"`
}
