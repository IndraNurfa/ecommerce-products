package api

import (
	"ecommerce-products/constants"
	"ecommerce-products/helpers"
	"ecommerce-products/internal/interfaces"
	"ecommerce-products/internal/models"
	"net/http"

	"github.com/labstack/echo/v5"
)

type ProductAPI struct {
	ProductService interfaces.IProductService
}

func (api *ProductAPI) CreateProduct(e *echo.Context) error {
	var (
		log = helpers.Logger
	)
	req := &models.Product{}

	if err := e.Bind(&req); err != nil {
		log.Error("failed to parse request: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
	}

	if err := req.Validate(); err != nil {
		log.Error("failed to validate request: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
	}

	resp, err := api.ProductService.CreateProduct(e.Request().Context(), req)
	if err != nil {
		log.Error("failed to create product : ", err)
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, constants.ErrFailedBadRequest, nil)
	}

	return helpers.SendResponseHTTP(e, http.StatusOK, constants.SuccessMessage, resp)
}
