package cmd

import (
	"ecommerce-products/external"
	"ecommerce-products/helpers"
	"ecommerce-products/internal/api"
	"ecommerce-products/internal/interfaces"
	"ecommerce-products/internal/repository"
	"ecommerce-products/internal/services"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func ServeHTTP() {
	d := dependencyInject()
	e := echo.New()

	e.Use(middleware.RequestLogger())
	e.GET("/health", d.HealthAPI.Health)

	v1 := e.Group("api/product/v1")

	v1.POST("", d.ProductAPI.CreateProduct, d.MiddlewareValidateAuth)

	categoryV1 := v1.Group("/category")
	categoryV1.POST("", d.CategoryAPI.CreateCategory, d.MiddlewareValidateAuth)

	if err := e.Start(":" + helpers.GetEnv("PORT", "9000")); err != nil {
		helpers.Logger.Error("failed to start server: ", err)
	}
}

type Dependency struct {
	HealthAPI *api.HealthAPI

	External interfaces.IExternal

	ProductAPI  interfaces.IProductAPI
	CategoryAPI interfaces.ICategoryAPI
}

func dependencyInject() Dependency {

	healthRepo := &repository.HealthRepository{
		DB: helpers.DB,
	}

	healthSvc := &services.HealthService{
		HealthRepo: healthRepo,
	}

	healthAPI := &api.HealthAPI{
		HealthService: healthSvc,
	}

	productRepo := &repository.ProductRepo{
		DB:    helpers.DB,
		Redis: helpers.RedisClient,
	}

	productSvc := &services.ProductService{
		ProductRepo: productRepo,
	}

	productAPI := &api.ProductAPI{
		ProductService: productSvc,
	}

	categoryRepo := &repository.CategoryRepo{
		DB: helpers.DB,
	}

	categorySvc := &services.CategoryService{
		CategoryRepo: categoryRepo,
	}

	categoryAPI := &api.CategoryAPI{
		CategoryService: categorySvc,
	}

	external := &external.External{}

	return Dependency{
		HealthAPI:   healthAPI,
		ProductAPI:  productAPI,
		CategoryAPI: categoryAPI,
		External:    external,
	}
}
