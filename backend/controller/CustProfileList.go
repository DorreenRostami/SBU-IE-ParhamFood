package controller

import (
	"net/http"

	model "github.com/DorreenRostami/IE_ParhamFood/model"
	"github.com/labstack/echo/v4"
)

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

func GetCustomerInfo(c echo.Context) error {
	var info CustomerMainInfo
	if err := c.Bind(&info); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	profiles := model.GetCustomerProfilesFromFile()
	var profInfo CustomerSignUpInfo
	for i := 0; i < len(profiles.Profiles); i++ {
		if profiles.Profiles[i].ID == info.ID {
			profInfo = CustomerSignUpInfo{
				Mobile:   profiles.Profiles[i].Mobile,
				Password: profiles.Profiles[i].Password,
				Name:     profiles.Profiles[i].Name,
				District: profiles.Profiles[i].District,
				Address:  profiles.Profiles[i].Address,
			}
			break
		}
	}
	return c.JSON(http.StatusOK, profInfo)
}
