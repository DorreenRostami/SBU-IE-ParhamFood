package restaurant

import (
	"net/http"

	fh "github.com/DorreenRostami/IE_ParhamFood/filehandler"
	"github.com/labstack/echo/v4"
)

type RestInfoReq struct {
	RID      int      `json:"restaurant_id"`
	RestInfo RestInfo `json:"restaurant_info"`
}

func restaurantExists(p1 fh.RestaurantProfile, p2 RestInfo) bool {
	return (p1.Name == p2.Name && p1.District == p2.District && p1.Address == p2.Address)
}

func UpdateInfo(c echo.Context) error { //needs every info field
	var info RestInfoReq
	if err := c.Bind(&info); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	profiles := fh.GetProfilesFromFile()
	for i := 0; i < len(profiles.Profiles); i++ {
		if profiles.Profiles[i].ID == info.RID {
			for j := 0; j < len(profiles.Profiles); j++ {
				if i == j {
					continue
				}
				if profiles.Profiles[j].Email == info.RestInfo.Email {
					return c.JSON(http.StatusConflict, ResponseMessage{
						StatusCode: http.StatusBadRequest,
						Message:    "This email has already been used.",
					})
				}
				if restaurantExists(profiles.Profiles[j], info.RestInfo) {
					return c.JSON(http.StatusConflict, ResponseMessage{
						StatusCode: http.StatusBadRequest,
						Message:    "This restaurant already exists.",
					})
				}
			}
			profiles.Profiles[i] = fh.RestaurantProfile{
				Email:       info.RestInfo.Email,
				Password:    info.RestInfo.Password,
				ID:          info.RID,
				Name:        info.RestInfo.Name,
				District:    info.RestInfo.District,
				Address:     info.RestInfo.Address,
				Open:        info.RestInfo.Open,
				Close:       info.RestInfo.Close,
				Dishes:      profiles.Profiles[i].Dishes,
				FixedCost:   info.RestInfo.FixedCost,
				FixedMinute: info.RestInfo.FixedMinute,
				Orders:      profiles.Profiles[i].Orders,
				Reviews:     profiles.Profiles[i].Reviews,
			}
			break
		}
		if i == len(profiles.Profiles)-1 {
			return c.JSON(http.StatusConflict, ResponseMessage{
				StatusCode: http.StatusBadRequest,
				Message:    "Wrong restaurant ID.",
			})
		}
	}
	fh.UpdateProfileFile(profiles)
	return c.JSON(http.StatusOK, info)
}
