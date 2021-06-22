package restaurantservices

import (
	"net/http"

	model "github.com/DorreenRostami/IE_ParhamFood/model"
	"github.com/labstack/echo/v4"
)


func GetOrders(c echo.Context) error {
	var id RestID
	if err := c.Bind(&id); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	ords := model.GetOrdersFromFile()
	var ans []model.Order
	for i := 0; i < len(ords.Orders); i++ {
		if ords.Orders[i].RestaurantID == id.RID {
			ans = append(ans, ords.Orders[i])
		}
	}
	return c.JSON(http.StatusOK, model.Orders{
		Orders: ans,
	})
}

func ChangeOrderStatus(c echo.Context) error { //needs order_id
	var order model.Order
	if err := c.Bind(&order); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	orders := model.GetOrdersFromFile()
	for i := 0; i < len(orders.Orders); i++ {
		if orders.Orders[i].OrderID == order.OrderID {
			orders.Orders[i].Status = orders.Orders[i].Status + 1
			order = orders.Orders[i]
			break
		}
		if i == len(orders.Orders)-1 {
			return c.JSON(http.StatusConflict, model.ResponseMessage{
				StatusCode: http.StatusBadRequest,
				Message:    "Wrong order ID.",
			})
		}
	}
	model.UpdateOrdersFile(orders)
	return c.JSON(http.StatusOK, order)
}
