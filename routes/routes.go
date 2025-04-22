package routes

import (
	"yubi-fullstack-test/handlers"
	"yubi-fullstack-test/repositories"
	"yubi-fullstack-test/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRoutes(db *gorm.DB, r *gin.Engine) {
	salesOrderRepository := repositories.NewSalesOrderRepository(db)
	salesOrderDetailRepository := repositories.NewSoDtRepository(db)
	salesOrderService := services.NewSalesOrderService(salesOrderRepository, salesOrderDetailRepository)
	salesOrderDetailService := services.NewSoDtService(salesOrderDetailRepository, salesOrderService, salesOrderRepository)
	salesOrderService.SetSoDtService(salesOrderDetailService)
	salesOrderhandler := handlers.NewSalesOrderHandler(salesOrderService)
	salesOrderDetailhandler := handlers.NewSoDtHandler(salesOrderDetailService)

	// API group for versioning (v1)
	api := r.Group("/api/v1")

	// Sales Orders Routes
	salesOrders := api.Group("/sales-orders")
	{

		// Get all sales orders
		salesOrders.GET("/", salesOrderhandler.GetAllSalesOrder)

		// Get one sales order by soId
		salesOrders.GET("/:soId", salesOrderhandler.GetSalesOrderById)

		// Create new sales order
		salesOrders.POST("/", salesOrderhandler.CreateSalesOrder)

		// Update sales order by soId
		salesOrders.PUT("/:soId", salesOrderhandler.UpdateSalesOrder)

		// Delete sales order by soId (also deletes all details)
		salesOrders.DELETE("/:soId", salesOrderhandler.DeleteSalesOrder)
	}

	// Sales Order Details Routes
	salesOrderDetails := api.Group("/sales-orders/:soId/details")
	{

		// Get all sales order details for soId
		salesOrderDetails.GET("/", salesOrderDetailhandler.GetAllBySalesOrderId)

		// Get one sales order detail by soDtId
		salesOrderDetails.GET("/:soDtId", salesOrderDetailhandler.GetBySalesOrderIdAndDetailId)

		// Create new sales order detail for soId
		salesOrderDetails.POST("/", salesOrderDetailhandler.CreateSoDt)

		// Update sales order detail by soDtId under soId
		salesOrderDetails.PUT("/:soDtId", salesOrderDetailhandler.UpdateSoDt)

		// Delete sales order detail by soDtId under soId
		salesOrderDetails.DELETE("/:soDtId", salesOrderDetailhandler.DeleteSoDt)
	}
}
