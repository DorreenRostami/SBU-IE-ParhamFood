package customerservices

import (
	"net/http"

	model "github.com/DorreenRostami/IE_ParhamFood/model"
	"github.com/labstack/echo/v4"
)

type CustInfoReq struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
	CID      int    `json:"id"`
	Name     string `json:"name"`
	District string `json:"district"`
	Address  string `json:"address"`
}

func UpdateCustomerInfo(c echo.Context) error { //needs every info field
	var info CustInfoReq
	if err := c.Bind(&info); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	profiles := model.GetCustomerProfilesFromFile()
	for i := 0; i < len(profiles.Profiles); i++ {
		if profiles.Profiles[i].ID == info.CID {
			for j := 0; j < len(profiles.Profiles); j++ {
				if i == j {
					continue
				}
				if profiles.Profiles[j].Mobile == info.Mobile {
					return c.JSON(http.StatusConflict, model.ResponseMessage{
						StatusCode: http.StatusBadRequest,
						Message:    "This mobile number has already been used.",
					})
				}
			}
			profiles.Profiles[i] = model.CustomerProfile{
				Mobile:   info.Mobile,
				Password: info.Password,
				ID:       info.CID,
				Name:     info.Name,
				District: info.District,
				Address:  info.Address,
				Balance:  profiles.Profiles[i].Balance,
			}
			break
		}
		if i == len(profiles.Profiles)-1 {
			return c.JSON(http.StatusConflict, model.ResponseMessage{
				StatusCode: http.StatusBadRequest,
				Message:    "Wrong customer ID.",
			})
		}
	}
	model.UpdateCustomerProfileFile(profiles)
	return c.JSON(http.StatusOK, info)
}
