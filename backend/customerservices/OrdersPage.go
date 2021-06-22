package customerservices

import (
	"net/http"

	model "github.com/DorreenRostami/IE_ParhamFood/model"
	"github.com/labstack/echo/v4"
)

type CustomerOrdersReq struct {
	ID int `json:"customer_id"`
}

func GetOrders(c echo.Context) error {
	var req CustomerOrdersReq
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	ords := model.GetOrdersFromFile()
	var ans []model.Order
	for i := 0; i < len(ords.Orders); i++ {
		if ords.Orders[i].CustomerID == req.ID {
			ans = append(ans, ords.Orders[i])
		}
	}
	return c.JSON(http.StatusOK, model.Orders{
		Orders: ans,
	})
}
