package dto

type SalesOrderDetailCreateRequest struct {
	RefType     *string  `json:"ref_type" binding:"required"`
	ItemType    *string  `json:"item_type" binding:"required"`
	ProductUUID *string  `json:"product_uuid" binding:"required"`
	ItemUnitID  *uint    `json:"item_unit_id" binding:"required"`
	PriceSell   *float64 `json:"price_sell" binding:"required"`
	Qty         *float64 `json:"qty" binding:"required"`
	DiscPerc    *float64 `json:"disc_perc" binding:"required"`
	DiscAm      *float64 `json:"disc_am" binding:"required"`
	Remark      *string  `json:"remark"`
}

type SalesOrderDetailUpdateRequest struct {
	RefType     *string  `json:"ref_type"`
	ItemType    *string  `json:"item_type"`
	ProductUUID *string  `json:"product_uuid"`
	ItemUnitID  *uint    `json:"item_unit_id"`
	PriceSell   *float64 `json:"price_sell"`
	Qty         *float64 `json:"qty"`
	DiscPerc    *float64 `json:"disc_perc"`
	DiscAm      *float64 `json:"disc_am"`
	Remark      *string  `json:"remark"`
}
