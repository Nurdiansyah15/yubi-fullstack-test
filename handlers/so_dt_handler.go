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

type SoDtHandler interface {
	GetAllBySalesOrderId(c *gin.Context)
	GetBySalesOrderIdAndDetailId(c *gin.Context)
	CreateSoDt(c *gin.Context)
	UpdateSoDt(c *gin.Context)
	DeleteSoDt(c *gin.Context)
}

type soDtHandler struct {
	soDtService services.SoDtService
}

func NewSoDtHandler(soDtService services.SoDtService) SoDtHandler {
	return &soDtHandler{soDtService: soDtService}
}

func (handler *soDtHandler) GetAllBySalesOrderId(c *gin.Context) {

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

	soDts, errCode := handler.soDtService.FindAllBySalesOrderId(uint(salesOrderIdUint))
	if errCode != "" {
		utils.ApiErrorResponse(c, 500, "INTERNAL_SERVER_ERROR", errCode)
		return
	}

	utils.ApiResponse(c, 200, "OK", soDts)
}

func (handler *soDtHandler) GetBySalesOrderIdAndDetailId(c *gin.Context) {
	salesOrderId := c.Param("soId")
	soDtId := c.Param("soDtId")

	if salesOrderId == "" || soDtId == "" {
		utils.ApiErrorResponse(c, 400, "BAD_REQUEST", "salesOrderId and soDtId are required")
		return
	}

	salesOrderIdUint64, err := strconv.ParseUint(salesOrderId, 10, 32)
	if err != nil {
		utils.ApiErrorResponse(c, 400, "BAD_REQUEST", "salesOrderId must be a number")
		return
	}
	soDtIdUint64, err := strconv.ParseUint(soDtId, 10, 32)
	if err != nil {
		utils.ApiErrorResponse(c, 400, "BAD_REQUEST", "soDtId must be a number")
		return
	}

	soDt, errCode := handler.soDtService.FindAllBySalesOrderIdAndDetailId(
		uint(salesOrderIdUint64),
		uint(soDtIdUint64),
	)

	if errCode != "" {
		switch errCode {
		case "SO_DT_NOT_FOUND_404":
			utils.ApiErrorResponse(c, 404, "SO_DT_NOT_FOUND_404", "Sales Order Detail not found")
		case "DATABASE_ERROR_500":
			utils.ApiErrorResponse(c, 500, "INTERNAL_SERVER_ERROR", "Database error")
		default:
			utils.ApiErrorResponse(c, 500, "INTERNAL_SERVER_ERROR", "Unknown error")
		}
		return
	}

	utils.ApiResponse(c, 200, "OK", soDt)
}

func (handler *soDtHandler) CreateSoDt(c *gin.Context) {

	salesOrderId := c.Param("soId")

	if salesOrderId == "" {
		utils.ApiErrorResponse(c, 400, "BAD_REQUEST", "salesOrderId is required")
		return
	}

	salesOrderIdUint64, err := strconv.ParseUint(salesOrderId, 10, 32)
	if err != nil {
		utils.ApiErrorResponse(c, 400, "BAD_REQUEST", "salesOrderId must be a number")
		return
	}

	var soDtCreateRequest dto.SalesOrderDetailCreateRequest

	if err := c.ShouldBindJSON(&soDtCreateRequest); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			ref := reflect.TypeOf(soDtCreateRequest)
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

	soDt := models.SoDt{
		SalesOrderID: uint(salesOrderIdUint64),
		RefType:      *soDtCreateRequest.RefType,
		ItemType:     *soDtCreateRequest.ItemType,
		ProductUUID:  *soDtCreateRequest.ProductUUID,
		ItemUnitID:   *soDtCreateRequest.ItemUnitID,
		PriceSell:    *soDtCreateRequest.PriceSell,
		Qty:          *soDtCreateRequest.Qty,
		DiscPerc:     *soDtCreateRequest.DiscPerc,
		DiscAm:       *soDtCreateRequest.DiscAm,
		Remark:       *soDtCreateRequest.Remark,
	}

	soDtStored, errCode := handler.soDtService.Store(soDt)
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

	utils.ApiResponse(c, 200, "OK", soDtStored)
}

func (handler *soDtHandler) UpdateSoDt(c *gin.Context) {
	salesOrderId := c.Param("soId")
	soDtId := c.Param("soDtId")

	if salesOrderId == "" || soDtId == "" {
		utils.ApiErrorResponse(c, 400, "BAD_REQUEST", "salesOrderId and soDtId are required")
		return
	}

	soDtIdUint, err := strconv.ParseUint(soDtId, 10, 32)
	if err != nil {
		utils.ApiErrorResponse(c, 400, "BAD_REQUEST", "soDtId must be a number")
		return
	}

	salesOrderIdUint64, err := strconv.ParseUint(salesOrderId, 10, 32)
	if err != nil {
		utils.ApiErrorResponse(c, 400, "BAD_REQUEST", "salesOrderId must be a number")
		return
	}

	var soDtUpdateRequest dto.SalesOrderDetailUpdateRequest

	if err := c.ShouldBindJSON(&soDtUpdateRequest); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			ref := reflect.TypeOf(soDtUpdateRequest)
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

	soDt := models.SoDt{
		ID:           uint(soDtIdUint),
		SalesOrderID: uint(salesOrderIdUint64),
		RefType:      *soDtUpdateRequest.RefType,
		ItemType:     *soDtUpdateRequest.ItemType,
		ProductUUID:  *soDtUpdateRequest.ProductUUID,
		ItemUnitID:   *soDtUpdateRequest.ItemUnitID,
		PriceSell:    *soDtUpdateRequest.PriceSell,
		Qty:          *soDtUpdateRequest.Qty,
		DiscPerc:     *soDtUpdateRequest.DiscPerc,
		DiscAm:       *soDtUpdateRequest.DiscAm,
		Remark:       *soDtUpdateRequest.Remark,
	}

	soDtUpdated, errCode := handler.soDtService.Update(soDt)
	if errCode != "" {
		switch errCode {
		case "SALES_ORDER_NOT_FOUND_404":
			utils.ApiErrorResponse(c, 404, "NOT_FOUND", errCode)
		case "SO_DT_NOT_FOUND_404":
			utils.ApiErrorResponse(c, 404, "SO_DT_NOT_FOUND_404", err)
		case "DATABASE_ERROR_500":
			utils.ApiErrorResponse(c, 500, "INTERNAL_SERVER_ERROR", err)
		default:
			utils.ApiErrorResponse(c, 500, "INTERNAL_SERVER_ERROR", err)
		}
		return
	}

	utils.ApiResponse(c, 200, "OK", soDtUpdated)
}

func (handler *soDtHandler) DeleteSoDt(c *gin.Context) {
	salesOrderId := c.Param("soId")
	soDtId := c.Param("soDtId")

	if salesOrderId == "" || soDtId == "" {
		utils.ApiErrorResponse(c, 400, "BAD_REQUEST", "salesOrderId and soDtId are required")
		return
	}

	soDtIdUint, err := strconv.ParseUint(soDtId, 10, 32)
	if err != nil {
		utils.ApiErrorResponse(c, 400, "BAD_REQUEST", "Invalid sales order detail ID")
		return
	}

	salesOrderIdUint, err := strconv.ParseUint(salesOrderId, 10, 32)
	if err != nil {
		utils.ApiErrorResponse(c, 400, "BAD_REQUEST", "salesOrderId must be a number")
		return
	}

	soDtDeleted, errCode := handler.soDtService.Delete(uint(salesOrderIdUint), uint(soDtIdUint))
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

	utils.ApiResponse(c, 200, "OK", soDtDeleted)
}
