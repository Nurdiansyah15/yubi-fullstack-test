package seeders

import (
	"time"
	"yubi-fullstack-test/database" // pastikan import package database jika ada
	"yubi-fullstack-test/models"
)

func SeedSalesOrders() {
	// Data dummy untuk SalesOrder
	salesOrders := []models.SalesOrder{
		{
			CustomerID:    1,
			OrderTypeID:   1,
			CurrencyID:    1,
			WarehouseID:   1,
			PaymentID:     1,
			VatID:         1,
			Pph23ID:       1,
			RevNo:         "REV001",
			PoBuyerNo:     "PO001",
			PoBuyerNoOri:  "POORI001",
			SalesOrderNo:  "SO001",
			Remark:        "First Order",
			ShipDest:      "Warehouse A",
			Status:        "Pending",
			ExchangeRate:  1.0,
			VatPerc:       10.0,
			Pph23Perc:     5.0,
			MarkupPerc:    20.0,
			IsVat:         true,
			IsPph23:       true,
			DiscAm:        100.0,
			DiscPerc:      10.0,
			DiscFinal:     90.0,
			TotalQty:      10,
			Subtotal:      1000.0,
			TotalDiscount: 100.0,
			TotalPph23:    50.0,
			TotalVat:      100.0,
			GrandTotal:    1050.0,
			QtyOut:        0,
			SaTotalAm:     1000.0,
			Total:         1050.0,
			OrderAt:       ptrTime(time.Now()),
			ShippingAt:    ptrTime(time.Now().Add(time.Hour * 24)),
			DueAt:         nil,
			ExpiredAt:     nil,
		},
	}

	for _, order := range salesOrders {
		database.DB.Create(&order)
	}
}

func ptrTime(t time.Time) *time.Time {
	return &t
}

func SeedSoDts() {
	soDts := []models.SoDt{
		{
			ProductUUID:     "product-uuid-001",
			SalesOrderID:    1,
			ItemUnitID:      1,
			VatID:           1,
			Pph23ID:         1,
			RefID:           1,
			RefJSON:         "{}",
			RefType:         "type1",
			ItemType:        "product",
			ItemJSON:        "{}",
			GenCode:         "genCode001",
			Remark:          "First Item",
			VatPercAm:       100.0,
			Pph23PercAm:     50.0,
			MarkupPercAm:    20.0,
			VatPerc:         10.0,
			Pph23Perc:       5.0,
			MarkupPerc:      20.0,
			IsVat:           true,
			IsPph23:         true,
			IsLockMarkup:    false,
			IsLockPriceSell: false,
			QtyOut:          0,
			Qty:             10,
			PriceSell:       100.0,
			PriceBuy:        90.0,
			SubtotalSell:    1000.0,
			SubtotalBuy:     900.0,
			DiscAm:          50.0,
			DiscPerc:        10.0,
			DiscPercNum:     10.0,
			DiscFinal:       45.0,
			Subtotal:        950.0,
			SaTotalAm:       900.0,
			Total:           950.0,
			TotalDp:         50.0,
		},
	}

	for _, soDt := range soDts {
		database.DB.Create(&soDt)
	}
}
