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

type RestMainInfos struct {
	AllRests []RestHomePageInfo `json:"all_restaurants"`
}

func GetAllRestaurants(c echo.Context) error {
	var allRests RestMainInfos
	profiles := model.GetProfilesFromFile()
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
