package customerservices

import (
	"net/http"
	"strings"

	model "github.com/DorreenRostami/IE_ParhamFood/model"
	rst "github.com/DorreenRostami/IE_ParhamFood/restaurantservices"
	"github.com/labstack/echo/v4"
)

type CustomerHomePageReq struct {
	ID         int    `json:"customer_id"`
	SearchText string `json:"search_text"`
}

type RestMenuInfo struct {
	RID         int          `json:"id"`
	Name        string       `json:"name"`
	District    string       `json:"district"`
	Address     string       `json:"address"`
	Open        int          `json:"open"`
	Close       int          `json:"close"`
	Dishes      []model.Dish `json:"dishes"`
	FixedCost   int          `json:"fixed_cost"`
	FixedMinute int          `json:"fixed_minute"`
}

func GetRestaurantsByFood(c echo.Context) error {
	var req CustomerHomePageReq
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var rests RestHomePageInfos
	profiles := model.GetRestaurantProfilesFromFile()
	for i := 0; i < len(profiles.Profiles); i++ {
		menu := profiles.Profiles[i].Dishes
		for j := 0; j < len(menu); j++ {
			if strings.EqualFold(menu[j].Name, req.SearchText) {
				rests.AllRests = append(rests.AllRests, RestHomePageInfo{
					RID:      profiles.Profiles[i].ID,
					Name:     profiles.Profiles[i].Name,
					District: profiles.Profiles[i].District,
					Address:  profiles.Profiles[i].Address,
				})
				break
			}
		}
	}
	return c.JSON(http.StatusOK, rests)
}

func GetRestaurantsByName(c echo.Context) error {
	var req CustomerHomePageReq
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var rests RestHomePageInfos
	profiles := model.GetRestaurantProfilesFromFile()
	for i := 0; i < len(profiles.Profiles); i++ {
		if strings.EqualFold(profiles.Profiles[i].Name, req.SearchText) {
			rests.AllRests = append(rests.AllRests, RestHomePageInfo{
				RID:      profiles.Profiles[i].ID,
				Name:     profiles.Profiles[i].Name,
				District: profiles.Profiles[i].District,
				Address:  profiles.Profiles[i].Address,
			})
		}
	}
	return c.JSON(http.StatusOK, rests)
}

func GetRestaurantsByDistrict(c echo.Context) error {
	var req CustomerHomePageReq
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var rests RestHomePageInfos
	profiles := model.GetRestaurantProfilesFromFile()
	for i := 0; i < len(profiles.Profiles); i++ {
		if strings.EqualFold(profiles.Profiles[i].District, req.SearchText) {
			rests.AllRests = append(rests.AllRests, RestHomePageInfo{
				RID:      profiles.Profiles[i].ID,
				Name:     profiles.Profiles[i].Name,
				District: profiles.Profiles[i].District,
				Address:  profiles.Profiles[i].Address,
			})
		}
	}
	return c.JSON(http.StatusOK, rests)
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
