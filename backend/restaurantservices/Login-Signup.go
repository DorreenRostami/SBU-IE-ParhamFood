package restaurantservices

import (
	"net/http"

	"github.com/labstack/echo/v4"

	model "github.com/DorreenRostami/IE_ParhamFood/model"
)

type AdminLoginInfo struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AdminSignupInfo struct {
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

func equalsRestaurant(p1 model.RestaurantProfile, p2 AdminSignupInfo) bool {
	return (p1.Name == p2.Name && p1.District == p2.District && p1.Address == p2.Address)
}

func LoginAdmin(c echo.Context) error {
	var loginInfo AdminLoginInfo
	if err := c.Bind(&loginInfo); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	profiles := model.GetRestaurantProfilesFromFile()
	for i := 0; i < len(profiles.Profiles); i++ {
		if profiles.Profiles[i].Email == loginInfo.Email && profiles.Profiles[i].Password == loginInfo.Password {
			return c.JSON(http.StatusOK, RestID{
				RID: profiles.Profiles[i].ID,
			})
		}
	}
	return c.JSON(http.StatusUnauthorized, model.ResponseMessage{
		StatusCode: http.StatusUnauthorized,
		Message:    "Wrong username or password.",
	})
}

func SignUpAdmin(c echo.Context) error {
	var signupInfo AdminSignupInfo
	if err := c.Bind(&signupInfo); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	profiles := model.GetRestaurantProfilesFromFile()

	for i := 0; i < len(profiles.Profiles); i++ {
		if profiles.Profiles[i].Email == signupInfo.Email {
			return c.JSON(http.StatusConflict, model.ResponseMessage{
				StatusCode: http.StatusConflict,
				Message:    "This email has already been used.",
			})
		}
		if equalsRestaurant(profiles.Profiles[i], signupInfo) {
			return c.JSON(http.StatusConflict, model.ResponseMessage{
				StatusCode: http.StatusConflict,
				Message:    "This restaurant already exists.",
			})
		}
	}

	newProfile := model.RestaurantProfile{
		Email:       signupInfo.Email,
		Password:    signupInfo.Password,
		ID:          len(profiles.Profiles),
		Name:        signupInfo.Name,
		District:    signupInfo.District,
		Address:     signupInfo.Address,
		Open:        signupInfo.Open,
		Close:       signupInfo.Close,
		Dishes:      []model.Dish{},
		FixedCost:   signupInfo.FixedCost,
		FixedMinute: signupInfo.FixedMinute,
		Orders:      []model.Order{},
		Reviews:     []model.Review{},
	}
	profiles.Profiles = append(profiles.Profiles, newProfile)
	model.UpdateRestaurantProfileFile(profiles)
	return c.JSON(http.StatusOK, RestID{
		RID: newProfile.ID,
	})
}
