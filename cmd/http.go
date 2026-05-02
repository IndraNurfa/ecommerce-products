package cmd

import (
	"ecommerce-products/helpers"
	"ecommerce-products/internal/api"
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

	if err := e.Start(":" + helpers.GetEnv("PORT", "9000")); err != nil {
		helpers.Logger.Error("failed to start server: ", err)
	}
}

type Dependency struct {
	HealthAPI *api.HealthAPI
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

	return Dependency{
		HealthAPI: healthAPI,
	}
}
