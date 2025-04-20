package dto

type SalesOrderDetailResponse struct {
	ID           uint    `json:"id"`
	SalesOrderID uint    `json:"sales_order_id"`
	RefType      string  `json:"ref_type"`
	ItemType     string  `json:"item_type"`
	ProductUUID  string  `json:"product_uuid"`
	ItemUnitID   uint    `json:"item_unit_id"`
	PriceSell    float64 `json:"price_sell"`
	Qty          float64 `json:"qty"`
	DiscPerc     float64 `json:"disc_perc"`
	DiscAm       float64 `json:"disc_am"`
	TotalAm      float64 `json:"total_am"`
	Remark       string  `json:"remark"`
}
