package customerservices

import (
	"net/http"

	model "github.com/DorreenRostami/IE_ParhamFood/model"
	rst "github.com/DorreenRostami/IE_ParhamFood/restaurantservices"

	"github.com/labstack/echo/v4"
)

type OrderReq struct {
	CustomerID   int              `json:"customer_id"`
	RestaurantID int              `json:"restaurant_id"`
	DisheInfos   []model.DishInfo `json:"dishes"`
	Price        int              `json:"price"`
}

func GetRestaurantMenu(c echo.Context) error {
	var req rst.RestID
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var restInfo RestMenuInfo
	profiles := model.GetRestaurantProfilesFromFile()
	for i := 0; i < len(profiles.Profiles); i++ {
		if profiles.Profiles[i].ID == req.RID {
			restInfo = RestMenuInfo{
				RID:         profiles.Profiles[i].ID,
				Name:        profiles.Profiles[i].Name,
				District:    profiles.Profiles[i].District,
				Address:     profiles.Profiles[i].Address,
				Open:        profiles.Profiles[i].Open,
				Close:       profiles.Profiles[i].Close,
				Dishes:      profiles.Profiles[i].Dishes,
				FixedCost:   profiles.Profiles[i].FixedCost,
				FixedMinute: profiles.Profiles[i].FixedMinute,
			}
			break
		}
	}
	return c.JSON(http.StatusOK, restInfo)
}

func CompleteOrder(c echo.Context) error {
	var req OrderReq
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	custs := model.GetCustomerProfilesFromFile()
	for i := 0; i < len(custs.Profiles); i++ {
		if custs.Profiles[i].ID == req.CustomerID {
			if custs.Profiles[i].Balance < req.Price {
				return c.JSON(http.StatusFailedDependency, model.ResponseMessage{
					StatusCode: http.StatusFailedDependency,
					Message:    "Insufficient balance.",
				})
			}
			custs.Profiles[i].Balance = custs.Profiles[i].Balance - req.Price
			model.UpdateCustomerProfileFile(custs)
			break
		}
	}

	ords := model.GetOrdersFromFile()
	neworder := model.Order{
		OrderID:      len(ords.Orders),
		CustomerID:   req.CustomerID,
		RestaurantID: req.RestaurantID,
		DisheInfos:   req.DisheInfos,
		Price:        req.Price,
		Status:       0,
	}
	ords.Orders = append(ords.Orders, neworder)
	model.UpdateOrdersFile(ords)
	return c.JSON(http.StatusOK, neworder)
}
