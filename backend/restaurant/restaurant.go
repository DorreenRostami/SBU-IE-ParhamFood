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

type DishReq struct {
	RestaurantID int    `json:"restaurant_id"`
	Name         string `json:"name"`
	Price        int    `json:"price"`
}

func equalsRestaurant(p1 fh.Profile, p2 SignupInfo) bool {
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
			response := profiles.Profiles[i]
			return c.JSON(http.StatusOK, response)
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

	newProfile := fh.Profile{
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
	return c.JSON(http.StatusOK, newProfile)
}

func getRestDishes(id int) []fh.Dish {
	profiles := fh.GetProfilesFromFile()
	for i := 0; i < len(profiles.Profiles); i++ {
		if profiles.Profiles[i].ID == id {
			return profiles.Profiles[i].Dishes
		}
	}
	return nil
}

func updateDishes(id int, d []fh.Dish) {
	profiles := fh.GetProfilesFromFile()
	for i := 0; i < len(profiles.Profiles); i++ {
		if profiles.Profiles[i].ID == id {
			profiles.Profiles[i].Dishes = d
			break
		}
	}
	fh.UpdateProfileFile(profiles)
}

func AddDish(c echo.Context) error {
	var dish DishReq
	if err := c.Bind(&dish); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	dishes := getRestDishes(dish.RestaurantID)
	for i := 0; i < len(dishes); i++ {
		if dishes[i].Name == dish.Name {
			return c.JSON(http.StatusConflict, ResponseMessage{
				StatusCode: http.StatusConflict,
				Message:    "A dish with this name already exists.",
			})
		}
	}

	newDish := fh.Dish{
		Name:      dish.Name,
		Price:     dish.Price,
		Available: true,
	}

	profiles := fh.GetProfilesFromFile()
	for i := 0; i < len(profiles.Profiles); i++ {
		if profiles.Profiles[i].ID == dish.RestaurantID {
			profiles.Profiles[i].Dishes = append(profiles.Profiles[i].Dishes, newDish)
			break
		}
	}
	fh.UpdateProfileFile(profiles)

	// var dddd fh.Profile
	// pfs := fh.GetProfilesFromFile()
	// for i := 0; i < len(pfs.Profiles); i++ {
	// 	if pfs.Profiles[i].ID == dish.RestaurantID {
	// 		dddd = pfs.Profiles[i]
	// 		break
	// 	}
	// }

	return c.JSON(http.StatusOK, newDish)
}

func DeleteDish(c echo.Context) error {
	var dish DishReq
	if err := c.Bind(&dish); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	dishes := getRestDishes(dish.RestaurantID)
	for i := 0; i < len(dishes); i++ {
		if dishes[i].Name == dish.Name {
			dishes = append(dishes[:i], dishes[i+1:]...)
			break
		}
	}

	updateDishes(dish.RestaurantID, dishes)

	return c.JSON(http.StatusOK, nil)
}
