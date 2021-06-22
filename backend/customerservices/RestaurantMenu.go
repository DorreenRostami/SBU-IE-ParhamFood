package customerservices

import (
	"net/http"
	"strconv"
	"time"

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

type OrderStatus struct {
	Status   int `json:"status"`
	MinsLeft int `json:"mins_left"`
}

func GetRestaurantMenu(c echo.Context) error {
	var req rst.RestID
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	profiles := model.GetRestaurantProfilesFromFile()
	for i := 0; i < len(profiles.Profiles); i++ {
		if profiles.Profiles[i].ID == req.RID {
			var dishes []model.Dish
			for j := 0; j < len(profiles.Profiles[i].Dishes); j++ {
				if profiles.Profiles[i].Dishes[j].Available {
					dishes = append(dishes, profiles.Profiles[i].Dishes[j])
				}
			}
			return c.JSON(http.StatusOK, RestMenuInfo{
				RID:         profiles.Profiles[i].ID,
				Name:        profiles.Profiles[i].Name,
				District:    profiles.Profiles[i].District,
				Address:     profiles.Profiles[i].Address,
				Open:        profiles.Profiles[i].Open,
				Close:       profiles.Profiles[i].Close,
				Dishes:      dishes,
				FixedCost:   profiles.Profiles[i].FixedCost,
				FixedMinute: profiles.Profiles[i].FixedMinute,
			})
		}
	}
	return c.JSON(http.StatusConflict, model.ResponseMessage{
		StatusCode: http.StatusConflict,
		Message:    "Wrong restaurant ID.",
	})
}

func CompleteOrder(c echo.Context) error {
	var req OrderReq
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var peik int
	rests := model.GetRestaurantProfilesFromFile()
	for i := 0; i < len(rests.Profiles); i++ {
		if rests.Profiles[i].ID == req.RestaurantID {
			peik = rests.Profiles[i].FixedCost
			break
		}
	}

	custs := model.GetCustomerProfilesFromFile()
	for i := 0; i < len(custs.Profiles); i++ {
		if custs.Profiles[i].ID == req.CustomerID {
			if custs.Profiles[i].Balance < req.Price+peik {
				return c.JSON(http.StatusFailedDependency, model.ResponseMessage{
					StatusCode: http.StatusFailedDependency,
					Message:    "Insufficient balance.",
				})
			}
			custs.Profiles[i].Balance = custs.Profiles[i].Balance - req.Price - peik
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
		TimeOfOrder:  time.Now(),
	}
	ords.Orders = append(ords.Orders, neworder)
	model.UpdateOrdersFile(ords)
	return c.JSON(http.StatusOK, neworder)
}

func GetOrderStatus(c echo.Context) error {
	number := c.Param("number")
	num, _ := strconv.Atoi(number)

	ords := model.GetOrdersFromFile()
	rests := model.GetRestaurantProfilesFromFile()
	for i := 0; i < len(ords.Orders); i++ {
		if ords.Orders[i].OrderID == num {
			var fixedMinute int
			for j := 0; j < len(rests.Profiles); j++ {
				if rests.Profiles[i].ID == ords.Orders[i].RestaurantID {
					fixedMinute = rests.Profiles[i].FixedMinute
					break
				}
			}
			minsPassed := int(time.Since(ords.Orders[i].TimeOfOrder).Minutes())
			minsLeft := 0
			if ords.Orders[i].Status != 4 && minsPassed < fixedMinute {
				minsLeft = fixedMinute - minsPassed
			}
			return c.JSON(http.StatusOK, OrderStatus{
				Status:   ords.Orders[i].Status,
				MinsLeft: minsLeft,
			})
		}
	}
	return c.JSON(http.StatusConflict, model.ResponseMessage{
		StatusCode: http.StatusConflict,
		Message:    "No such order exists.",
	})
}
