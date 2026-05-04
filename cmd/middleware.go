package cmd

import (
	"ecommerce-products/helpers"
	"net/http"

	"github.com/labstack/echo/v5"
)

func (d *Dependency) MiddlewareValidateAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) error {
		auth := c.Request().Header.Get("Authorization")
		if auth == "" {
			helpers.Logger.Errorf("authorization empty")
			return helpers.SendResponseHTTP(c, http.StatusUnauthorized, "unauthorized", nil)
		}

		profile, err := d.External.GetProfile(c.Request().Context(), auth)
		if err != nil {
			helpers.Logger.Errorf("authorization empty")
			return helpers.SendResponseHTTP(c, http.StatusUnauthorized, "unauthorized", nil)
		}

		c.Set("profile", profile)

		return next(c)
	}

}
