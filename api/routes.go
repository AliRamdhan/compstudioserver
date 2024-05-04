package api

import (
	"net/http"

	"github.com/AliRamdhan/compstudioserver/internal/handlers"
	"github.com/AliRamdhan/compstudioserver/internal/middlewares"
	"github.com/AliRamdhan/compstudioserver/internal/services"

	"github.com/gin-gonic/gin"
)

func ServiceAuth(r *gin.Engine, authUser *services.AuthService) {
	authHandler := handlers.NewAuthHandler(authUser)
	authRoutes := r.Group("/auth")
	{
		authRoutes.GET("/user/all", authHandler.GetAllUser)
		authRoutes.POST("/register", authHandler.RegisterAuth)
		authRoutes.POST("/login", authHandler.Login)
		//authRoutes.POST("/home", authHandler.Home).Use(middlewares.Auth())
		securedUserRoutes := r.Group("/home").Use(middlewares.UserAuth())
		{
			securedUserRoutes.GET("/user", authHandler.Home)
		}
		securedAdminRoutes := r.Group("/home").Use(middlewares.AdminAuth())
		{
			securedAdminRoutes.GET("/admin", authHandler.Home)
		}
	}
	// Handler OPTIONS untuk CORS preflight
	r.OPTIONS("/home/admin", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
}

func ServiceRole(r *gin.Engine, roleService *services.RoleService) {
	roleHandler := handlers.NewRoleHandler(roleService)
	roleRoutes := r.Group("/role")
	{
		roleRoutes.POST("/create", roleHandler.CreateRole)
		roleRoutes.GET("/all", roleHandler.GetAllRole)
		roleRoutes.PUT("/update/:roleId", roleHandler.UpdateRole)
		roleRoutes.DELETE("delete/:roleId", roleHandler.DeleteRole)
	}
}

func ServiceTrackStatus(r *gin.Engine, roleTrackStatus *services.TrackStatusService) {
	trackStatusHandler := handlers.NewTrackStatusHandler(roleTrackStatus)
	trackStatusRoutes := r.Group("/track-status")
	{
		trackStatusRoutes.POST("/create", trackStatusHandler.CreateTrackStatus)
		trackStatusRoutes.GET("/all", trackStatusHandler.GetAllTrackStatus)
		trackStatusRoutes.PUT("/update/:statusId", trackStatusHandler.UpdateTrackStatus)
		trackStatusRoutes.DELETE("delete/:statusId", trackStatusHandler.DeleteTrackStatus)
	}
}

func ServiceCategory(r *gin.Engine, categoryService *services.CategoryService) {
	routesHandler := handlers.NewServiceCategoryHandler(categoryService)
	routesService := r.Group("/service-category")
	{
		routesService.POST("/create", routesHandler.CreateServiceCategory)
		routesService.GET("/all", routesHandler.GetAllServiceCategory)
		routesService.PUT("/update/:catId", routesHandler.UpdateServiceCategory)
		routesService.DELETE("delete/:catId", routesHandler.DeleteServiceCategory)
	}
}

func Servicecomp(r *gin.Engine, serviceComp *services.ServiceComp) {
	routesHandler := handlers.NewServiceCompHandler(serviceComp)
	routesService := r.Group("/service")
	{
		routesService.POST("/create", routesHandler.CreateserviceComp)
		routesService.GET("/all", routesHandler.GetAllService)
		routesService.GET("/:serviceId", routesHandler.GetDetailServiceById)
		routesService.PUT("/update/:serviceId", routesHandler.UpdateService)
		routesService.DELETE("delete/:serviceId", routesHandler.DeleteService)
	}
}

func ServiceTrackComp(r *gin.Engine, serviceComp *services.TrackService) {
	routesHandler := handlers.NewTrackHandler(serviceComp)
	routesService := r.Group("/track")
	{
		routesService.POST("/create", routesHandler.CreatetrackService)
		routesService.POST("/create/:trackNumber", routesHandler.CreateProgressTrackStatusByTrackNumber)
		routesService.GET("/all", routesHandler.GetAllTrack)
		routesService.GET("/:trackNumber", routesHandler.GetTrackStatusByTrackNumber)
		routesService.PUT("/update/:trackId", routesHandler.UpdateTrack)
		routesService.DELETE("delete/:trackId", routesHandler.DeleteTrack)
	}
}

func ServiceProducts(r *gin.Engine, productService *services.ProductServices) {
	productHandler := handlers.NewProductHandler(productService)
	// Product routes
	productRoutes := r.Group("/products")
	{
		productRoutes.POST("/create", middlewares.ValidateFileMiddleware(), productHandler.CreateProduct)
		productRoutes.GET("/all", productHandler.GetAllProducts)
		productRoutes.PUT("/update/:productId", productHandler.UpdateProduct)
		productRoutes.DELETE("/delete/:productId", productHandler.DeleteProduct)
	}
}

func MessageService(r *gin.Engine, messageService *services.MessageService) {
	routesHandler := handlers.NewMessageHandler(messageService)
	routesService := r.Group("/message")
	{
		routesService.POST("/create", routesHandler.CreateMessage)
		// routesService.PUT("/read/:serviceId", routesHandler.MarkMessageAsRead)
		routesService.GET("/all", routesHandler.GetAllMessage)
		routesService.GET("service/:serviceId/message", routesHandler.GetMessageAndMarkAsRead)
		routesService.PUT("/update/:messageId", routesHandler.UpdateMessage)
		routesService.DELETE("delete/:messageId", routesHandler.DeleteMessage)
	}
}
