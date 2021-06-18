package restaurant

import (
	"net/http"

	fh "github.com/DorreenRostami/IE_ParhamFood/filehandler"
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
	ID int `json:"id"`
}

type Menu struct {
	Dishes []fh.Dish `json:"dishes"`
}

type Orders struct {
	Orders []fh.Order `json:"orders"`
}

type Reviews struct {
	Reviews []fh.Review `json:"reviews"`
}

func GetMenu(c echo.Context) error {
	var id RestID
	if err := c.Bind(&id); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, Menu{
		Dishes: getDishes(id.ID),
	})
}

func GetOrders(c echo.Context) error {
	var id RestID
	if err := c.Bind(&id); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	profiles := fh.GetProfilesFromFile()
	var orders []fh.Order
	for i := 0; i < len(profiles.Profiles); i++ {
		if profiles.Profiles[i].ID == id.ID {
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
	profiles := fh.GetProfilesFromFile()
	var rev []fh.Review
	for i := 0; i < len(profiles.Profiles); i++ {
		if profiles.Profiles[i].ID == id.ID {
			rev = profiles.Profiles[i].Reviews
			break
		}
	}
	return c.JSON(http.StatusOK, Reviews{
		Reviews: rev,
	})
}

func GetInfo(c echo.Context) error {
	var id RestID
	if err := c.Bind(&id); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	profiles := fh.GetProfilesFromFile()
	var info RestInfo
	for i := 0; i < len(profiles.Profiles); i++ {
		if profiles.Profiles[i].ID == id.ID {
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
