package controller

import (
	"net/http"

	model "github.com/DorreenRostami/IE_ParhamFood/model"
	"github.com/labstack/echo/v4"
)

type RestInfo struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	Name        string `json:"name"`
	District    string `json:"district"`
	Address     string `json:"address"`
	Open        int    `json:"open"`
	Close       int    `json:"close"`
	FixedCost   int    `json:"fixed_cost"`
	FixedMinute int    `json:"fixed_minute"`
}

type RestID struct {
	RID int `json:"restaurant_id"`
}

type Menu struct {
	Dishes []model.Dish `json:"dishe"`
}

type Orders struct {
	Orders []model.Order `json:"order"`
}

type Reviews struct {
	Reviews []model.Review `json:"review"`
}

func GetMenu(c echo.Context) error {
	var id RestID
	if err := c.Bind(&id); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, Menu{
		Dishes: getDishes(id.RID),
	})
}

func GetOrders(c echo.Context) error {
	var id RestID
	if err := c.Bind(&id); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	profiles := model.GetRestaurantProfilesFromFile()
	var orders []model.Order
	for i := 0; i < len(profiles.Profiles); i++ {
		if profiles.Profiles[i].ID == id.RID {
			orders = profiles.Profiles[i].Orders
			break
		}
	}
	return c.JSON(http.StatusOK, Orders{
		Orders: orders,
	})
}

func GetReviews(c echo.Context) error {
	var id RestID
	if err := c.Bind(&id); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	profiles := model.GetRestaurantProfilesFromFile()
	var rev []model.Review
	for i := 0; i < len(profiles.Profiles); i++ {
		if profiles.Profiles[i].ID == id.RID {
			rev = profiles.Profiles[i].Reviews
			break
		}
	}
	return c.JSON(http.StatusOK, Reviews{
		Reviews: rev,
	})
}

func GetRestaurantInfo(c echo.Context) error {
	var id RestID
	if err := c.Bind(&id); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	profiles := model.GetRestaurantProfilesFromFile()
	var info RestInfo
	for i := 0; i < len(profiles.Profiles); i++ {
		if profiles.Profiles[i].ID == id.RID {
			info = RestInfo{
				Email:       profiles.Profiles[i].Email,
				Password:    profiles.Profiles[i].Password,
				Name:        profiles.Profiles[i].Name,
				District:    profiles.Profiles[i].District,
				Address:     profiles.Profiles[i].Address,
				Open:        profiles.Profiles[i].Open,
				Close:       profiles.Profiles[i].Close,
				FixedCost:   profiles.Profiles[i].FixedCost,
				FixedMinute: profiles.Profiles[i].FixedMinute,
			}
			break
		}
	}
	return c.JSON(http.StatusOK, info)
}
