package server

import (
	"net/http"

	"github.com/kliment2000/go_kafka_pg/cache"
	"github.com/kliment2000/go_kafka_pg/database"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Start() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet},
	}))

	e.GET("/order/:order_uid", func(c echo.Context) error {
		orderUID := c.Param("order_uid")

		if order, ok := cache.Cache.Get(orderUID); ok {
			return c.JSON(http.StatusOK, order)
		}
		if order, err := database.GetOrder(orderUID); err == nil {
			return c.JSON(http.StatusOK, order.Data)
		}

		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "order not found",
		})
	})

	e.Logger.Fatal(e.Start(":8081"))
}
