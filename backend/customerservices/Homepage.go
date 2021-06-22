package customerservices

import (
	"net/http"
	"strings"

	model "github.com/DorreenRostami/IE_ParhamFood/model"

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

type RestHomePageInfo struct {
	RID      int    `json:"restaurant_id"`
	Name     string `json:"name"`
	District string `json:"district"`
	Address  string `json:"address"`
}

type RestHomePageInfos struct {
	AllRests []RestHomePageInfo `json:"restaurant_homepage_infos"`
}

func GetAllRestaurants(c echo.Context) error {
	var allRests RestHomePageInfos
	profiles := model.GetRestaurantProfilesFromFile()
	for i := 0; i < len(profiles.Profiles); i++ {
		allRests.AllRests = append(allRests.AllRests, RestHomePageInfo{
			RID:      profiles.Profiles[i].ID,
			Name:     profiles.Profiles[i].Name,
			District: profiles.Profiles[i].District,
			Address:  profiles.Profiles[i].Address,
		})
	}
	return c.JSON(http.StatusOK, allRests)
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
