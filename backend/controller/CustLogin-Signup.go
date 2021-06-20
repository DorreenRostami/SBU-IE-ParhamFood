package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	ptime "github.com/yaa110/go-persian-calendar"

	model "github.com/DorreenRostami/IE_ParhamFood/model"
)

type CustomerLoginInfo struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

type CustomerSignUpInfo struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
	Name     string `json:"name"`
	District string `json:"district"`
	Address  string `json:"address"`
}

type CustomerMainInfo struct {
	ID      int `json:"id"`
	Balance int `json:"balance"`
}

func LoginCustomer(c echo.Context) error {
	var loginInfo CustomerLoginInfo
	if err := c.Bind(&loginInfo); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	profiles := model.GetCustomerProfilesFromFile()
	for i := 0; i < len(profiles.Profiles); i++ {
		if profiles.Profiles[i].Mobile == loginInfo.Mobile && profiles.Profiles[i].Password == loginInfo.Password {
			return c.JSON(http.StatusOK, CustomerMainInfo{
				ID:      profiles.Profiles[i].ID,
				Balance: profiles.Profiles[i].Balance,
			})
		}
	}
	return c.JSON(http.StatusUnauthorized, ResponseMessage{
		StatusCode: http.StatusUnauthorized,
		Message:    "Wrong username or password.",
	})
}

func SignUpCustomer(c echo.Context) error {
	var signupInfo CustomerSignUpInfo
	if err := c.Bind(&signupInfo); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	profiles := model.GetCustomerProfilesFromFile()

	for i := 0; i < len(profiles.Profiles); i++ {
		if profiles.Profiles[i].Mobile == signupInfo.Mobile {
			return c.JSON(http.StatusConflict, ResponseMessage{
				StatusCode: http.StatusConflict,
				Message:    "This mobile number has already been used.",
			})
		}
	}

	bal := 0
	if ptime.Now().Year() < 1401 {
		bal = 1000000
	}

	newProfile := model.CustomerProfile{
		Mobile:   signupInfo.Mobile,
		Password: signupInfo.Password,
		ID:       len(profiles.Profiles) + 1,
		Name:     signupInfo.Name,
		District: signupInfo.District,
		Address:  signupInfo.Address,
		Balance:  bal,
	}

	profiles.Profiles = append(profiles.Profiles, newProfile)
	model.UpdateCustomerProfileFile(profiles)
	return c.JSON(http.StatusOK, CustomerMainInfo{
		ID:      newProfile.ID,
		Balance: newProfile.Balance,
	})
}
