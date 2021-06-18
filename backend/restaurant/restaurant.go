package restaurant

import (
	"net/http"

	"github.com/labstack/echo/v4"

	fh "github.com/DorreenRostami/IE_ParhamFood/filehandler"
)

type ResponseMessage struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func equalsRestaurant(p1 fh.Profile, p2 fh.Profile) bool {
	return (p1.Name == p2.Name && p1.District == p2.District && p1.Address == p2.Address)
}

func LoginAdmin(c echo.Context) error {
	var loginInfo fh.LoginInfo
	if err := c.Bind(&loginInfo); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	profiles := fh.GetProfilesFromFile()
	for i := 0; i < len(profiles.Profiles); i++ {
		if profiles.Profiles[i].Email == loginInfo.Email && profiles.Profiles[i].Password == loginInfo.Password {
			response := profiles.Profiles[i]
			return c.JSON(http.StatusOK, response)
		}
	}
	return c.JSON(http.StatusUnauthorized, ResponseMessage{StatusCode: http.StatusUnauthorized, Message: "Wrong username or password."})
}

func SignUpAdmin(c echo.Context) error {
	var signupInfo fh.Profile
	if err := c.Bind(&signupInfo); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	profiles := fh.GetProfilesFromFile()

	for i := 0; i < len(profiles.Profiles); i++ {
		if profiles.Profiles[i].Email == signupInfo.Email {
			return c.JSON(http.StatusConflict, ResponseMessage{StatusCode: http.StatusConflict, Message: "This email has already been used."})
		}
		if equalsRestaurant(profiles.Profiles[i], signupInfo) {
			return c.JSON(http.StatusConflict, ResponseMessage{StatusCode: http.StatusConflict, Message: "This restaurant already exists."})
		}
	}

	newProfile := fh.Profile{
		Email:     signupInfo.Email,
		Password:  signupInfo.Password,
		ID:        len(profiles.Profiles),
		Name:      signupInfo.Name,
		District:  signupInfo.District,
		Address:   signupInfo.Address,
		Open:      0,
		Close:     0,
		Dishes:    []fh.Dish{},
		FixedCost: 0,
		FixedTime: 0,
		Orders:    []fh.Order{},
		Reviews:   []fh.Review{},
	}
	fh.AddProfileToFile(profiles, newProfile)
	return c.JSON(http.StatusOK, newProfile)
}
