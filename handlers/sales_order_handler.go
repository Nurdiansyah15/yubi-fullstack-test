package handlers

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"yubi-fullstack-test/dto"
	"yubi-fullstack-test/models"
	"yubi-fullstack-test/services"
	"yubi-fullstack-test/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type SalesOrderHandler interface {
	GetAllSalesOrder(c *gin.Context)
	GetSalesOrderById(c *gin.Context)
	CreateSalesOrder(c *gin.Context)
	UpdateSalesOrder(c *gin.Context)
	DeleteSalesOrder(c *gin.Context)
}

type salesOrderHandler struct {
	salesOrderService services.SalesOrderService
}

func NewSalesOrderHandler(salesOrderService services.SalesOrderService) SalesOrderHandler {
	return &salesOrderHandler{salesOrderService: salesOrderService}
}

func (handler *salesOrderHandler) GetAllSalesOrder(c *gin.Context) {
	salesOrders, err := handler.salesOrderService.FindAll()
	if err != "" {
		utils.ApiErrorResponse(c, 500, "INTERNAL_SERVER_ERROR", err)
		return
	}

	utils.ApiResponse(c, 200, "OK", salesOrders)
}

func (handler *salesOrderHandler) GetSalesOrderById(c *gin.Context) {
	salesOrderId := c.Param("soId")

	if salesOrderId == "" {
		utils.ApiErrorResponse(c, 400, "BAD_REQUEST", "salesOrderId is required")
		return
	}

	salesOrderIdUint, err := strconv.ParseUint(salesOrderId, 10, 32)
	if err != nil {
		utils.ApiErrorResponse(c, 400, "BAD_REQUEST", "salesOrderId must be a number")
		return
	}

	salesOrder, errCode := handler.salesOrderService.FindById(uint(salesOrderIdUint))
	if errCode != "" {
		switch errCode {
		case "SALES_ORDER_NOT_FOUND_404":
			utils.ApiErrorResponse(c, 404, "SALES_ORDER_NOT_FOUND_404", err)
		case "DATABASE_ERROR_500":
			utils.ApiErrorResponse(c, 500, "INTERNAL_SERVER_ERROR", err)
		default:
			utils.ApiErrorResponse(c, 500, "INTERNAL_SERVER_ERROR", err)
		}
		return
	}

	utils.ApiResponse(c, 200, "OK", salesOrder)

}

func (handler *salesOrderHandler) CreateSalesOrder(c *gin.Context) {
	var salesOrderCreateRequest dto.SalesOrderCreateRequest

	if err := c.ShouldBindJSON(&salesOrderCreateRequest); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			ref := reflect.TypeOf(salesOrderCreateRequest)
			if ref.Kind() == reflect.Ptr {
				ref = ref.Elem()
			}

			var errorsList []string
			for _, fe := range ve {
				field, _ := ref.FieldByName(fe.Field())
				jsonTag := field.Tag.Get("json")
				if jsonTag == "" {
					jsonTag = strings.ToLower(fe.Field())
				}

				errorsList = append(errorsList, fmt.Sprintf("Field '%s' is %s", jsonTag, fe.Tag()))
			}

			utils.ApiErrorResponse(c, 400, "VALIDATION_ERROR", strings.Join(errorsList, "; "))
			return
		}

		utils.ApiErrorResponse(c, 400, "BAD_REQUEST", err.Error())
		return
	}

	salesOrder := models.SalesOrder{
		PoBuyerNo:     *salesOrderCreateRequest.PoBuyerNo,
		OrderTypeID:   *salesOrderCreateRequest.OrderTypeID,
		CustomerID:    *salesOrderCreateRequest.CustomerID,
		Status:        *salesOrderCreateRequest.Status,
		OrderAt:       salesOrderCreateRequest.OrderAt,
		ShippingAt:    salesOrderCreateRequest.ShippingAt,
		CurrencyID:    *salesOrderCreateRequest.CurrencyID,
		ExchangeRate:  *salesOrderCreateRequest.ExchangeRate,
		IsVat:         *salesOrderCreateRequest.IsVat,
		IsPph23:       *salesOrderCreateRequest.IsPph23,
		VatID:         *salesOrderCreateRequest.VatID,
		Pph23ID:       *salesOrderCreateRequest.Pph23ID,
		Subtotal:      *salesOrderCreateRequest.Subtotal,
		TotalDiscount: *salesOrderCreateRequest.TotalDiscount,
		TotalVat:      *salesOrderCreateRequest.TotalVat,
		TotalPph23:    *salesOrderCreateRequest.TotalPph23,
		GrandTotal:    *salesOrderCreateRequest.GrandTotal,
		Remark:        *salesOrderCreateRequest.Remark,
	}

	salesOrderCreated, errCode := handler.salesOrderService.Store(salesOrder)

	if errCode != "" {
		switch errCode {
		case "DATABASE_ERROR_500":
			utils.ApiErrorResponse(c, 500, "INTERNAL_SERVER_ERROR", errCode)
		default:
			utils.ApiErrorResponse(c, 500, "INTERNAL_SERVER_ERROR", errCode)
		}
		return
	}

	utils.ApiResponse(c, 200, "OK", salesOrderCreated)
}

func (handler *salesOrderHandler) UpdateSalesOrder(c *gin.Context) {
	salesOrderId := c.Param("soId")

	if salesOrderId == "" {
		utils.ApiErrorResponse(c, 400, "BAD_REQUEST", "salesOrderId is required")
		return
	}

	idUint, err := strconv.ParseUint(salesOrderId, 10, 32)
	if err != nil {
		utils.ApiErrorResponse(c, 400, "BAD_REQUEST", "Invalid sales order ID")
		return
	}

	var salesOrderUpdateRequest dto.SalesOrderUpdateRequest

	if err := c.ShouldBindJSON(&salesOrderUpdateRequest); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			ref := reflect.TypeOf(salesOrderUpdateRequest)
			if ref.Kind() == reflect.Ptr {
				ref = ref.Elem()
			}

			var errorsList []string
			for _, fe := range ve {
				field, _ := ref.FieldByName(fe.Field())
				jsonTag := field.Tag.Get("json")
				if jsonTag == "" {
					jsonTag = strings.ToLower(fe.Field())
				}

				errorsList = append(errorsList, fmt.Sprintf("Field '%s' is %s", jsonTag, fe.Tag()))
			}

			utils.ApiErrorResponse(c, 400, "VALIDATION_ERROR", strings.Join(errorsList, "; "))
			return
		}

		utils.ApiErrorResponse(c, 400, "BAD_REQUEST", err.Error())
		return
	}

	salesOrder := models.SalesOrder{
		ID:            uint(idUint),
		PoBuyerNo:     *salesOrderUpdateRequest.PoBuyerNo,
		OrderTypeID:   *salesOrderUpdateRequest.OrderTypeID,
		CustomerID:    *salesOrderUpdateRequest.CustomerID,
		Status:        *salesOrderUpdateRequest.Status,
		OrderAt:       salesOrderUpdateRequest.OrderAt,
		ShippingAt:    salesOrderUpdateRequest.ShippingAt,
		CurrencyID:    *salesOrderUpdateRequest.CurrencyID,
		ExchangeRate:  *salesOrderUpdateRequest.ExchangeRate,
		IsVat:         *salesOrderUpdateRequest.IsVat,
		IsPph23:       *salesOrderUpdateRequest.IsPph23,
		VatID:         *salesOrderUpdateRequest.VatID,
		Pph23ID:       *salesOrderUpdateRequest.Pph23ID,
		Subtotal:      *salesOrderUpdateRequest.Subtotal,
		TotalDiscount: *salesOrderUpdateRequest.TotalDiscount,
		TotalVat:      *salesOrderUpdateRequest.TotalVat,
		TotalPph23:    *salesOrderUpdateRequest.TotalPph23,
		GrandTotal:    *salesOrderUpdateRequest.GrandTotal,
		Remark:        *salesOrderUpdateRequest.Remark,
	}

	salesOrderUpdated, errCode := handler.salesOrderService.Update(salesOrder)

	if errCode != "" {
		switch errCode {
		case "DATABASE_ERROR_500":
			utils.ApiErrorResponse(c, 500, "INTERNAL_SERVER_ERROR", errCode)
		default:
			utils.ApiErrorResponse(c, 500, "INTERNAL_SERVER_ERROR", errCode)
		}
		return
	}

	utils.ApiResponse(c, 200, "OK", salesOrderUpdated)

}

func (handler *salesOrderHandler) DeleteSalesOrder(c *gin.Context) {
	salesOrderId := c.Param("soId")

	if salesOrderId == "" {
		utils.ApiErrorResponse(c, 400, "BAD_REQUEST", "salesOrderId are required")
		return
	}

	idUint, err := strconv.ParseUint(salesOrderId, 10, 32)
	if err != nil {
		utils.ApiErrorResponse(c, 400, "BAD_REQUEST", "Invalid sales order ID")
		return
	}

	salesOrderDeleted, errCode := handler.salesOrderService.Delete(uint(idUint))

	if errCode != "" {
		switch errCode {
		case "SALES_ORDER_NOT_FOUND_404":
			utils.ApiErrorResponse(c, 404, "NOT_FOUND", errCode)
		case "DATABASE_ERROR_500":
			utils.ApiErrorResponse(c, 500, "INTERNAL_SERVER_ERROR", errCode)
		default:
			utils.ApiErrorResponse(c, 500, "INTERNAL_SERVER_ERROR", errCode)
		}
		return
	}

	utils.ApiResponse(c, 200, "OK", salesOrderDeleted)

}
