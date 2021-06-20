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

type LoginInfo struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignupInfo struct {
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

func equalsRestaurant(p1 fh.RestaurantProfile, p2 SignupInfo) bool {
	return (p1.Name == p2.Name && p1.District == p2.District && p1.Address == p2.Address)
}

func LoginAdmin(c echo.Context) error {
	var loginInfo LoginInfo
	if err := c.Bind(&loginInfo); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	profiles := fh.GetProfilesFromFile()
	for i := 0; i < len(profiles.Profiles); i++ {
		if profiles.Profiles[i].Email == loginInfo.Email && profiles.Profiles[i].Password == loginInfo.Password {
			return c.JSON(http.StatusOK, RestID{
				RID: profiles.Profiles[i].ID,
			})
		}
	}
	return c.JSON(http.StatusUnauthorized, ResponseMessage{
		StatusCode: http.StatusUnauthorized,
		Message:    "Wrong username or password.",
	})
}

func SignUpAdmin(c echo.Context) error {
	var signupInfo SignupInfo
	if err := c.Bind(&signupInfo); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	profiles := fh.GetProfilesFromFile()

	for i := 0; i < len(profiles.Profiles); i++ {
		if profiles.Profiles[i].Email == signupInfo.Email {
			return c.JSON(http.StatusConflict, ResponseMessage{
				StatusCode: http.StatusConflict,
				Message:    "This email has already been used.",
			})
		}
		if equalsRestaurant(profiles.Profiles[i], signupInfo) {
			return c.JSON(http.StatusConflict, ResponseMessage{
				StatusCode: http.StatusConflict,
				Message:    "This restaurant already exists.",
			})
		}
	}

	newProfile := fh.RestaurantProfile{
		Email:       signupInfo.Email,
		Password:    signupInfo.Password,
		ID:          len(profiles.Profiles),
		Name:        signupInfo.Name,
		District:    signupInfo.District,
		Address:     signupInfo.Address,
		Open:        signupInfo.Open,
		Close:       signupInfo.Close,
		Dishes:      []fh.Dish{},
		FixedCost:   signupInfo.FixedCost,
		FixedMinute: signupInfo.FixedMinute,
		Orders:      []fh.Order{},
		Reviews:     []fh.Review{},
	}
	profiles.Profiles = append(profiles.Profiles, newProfile)
	fh.UpdateProfileFile(profiles)
	return c.JSON(http.StatusOK, RestID{
		RID: newProfile.ID,
	})
}
