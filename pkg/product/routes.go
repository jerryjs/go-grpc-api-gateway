package product

import (
	"github.com/gin-gonic/gin"
	"github.com/jerryjs/go-grpc-api-gateway/pkg/auth"
	"github.com/jerryjs/go-grpc-api-gateway/pkg/config"
	"github.com/jerryjs/go-grpc-api-gateway/pkg/product/routes"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/product")
	routes.Use(a.AuthRequired)
	routes.POST("/", svc.CreateProduct)
	routes.GET("/:id", svc.FindOne)
	routes.POST("/:id/stock/:stock", svc.AddStock)
}

func (svc *ServiceClient) AddStock(ctx *gin.Context) {
	routes.AddStock(ctx, svc.Client)
}

func (svc *ServiceClient) FindOne(ctx *gin.Context) {
	routes.FindOne(ctx, svc.Client)
}

func (svc *ServiceClient) CreateProduct(ctx *gin.Context) {
	routes.CreateProduct(ctx, svc.Client)
}
