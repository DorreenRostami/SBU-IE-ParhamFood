package controller

import (
	"fmt"
	"net/http"

	model "github.com/DorreenRostami/IE_ParhamFood/model"
	"github.com/labstack/echo/v4"
)

type DishReq struct {
	RestaurantID int    `json:"restaurant_id"`
	Name         string `json:"name"`
	Price        int    `json:"price"`
	Available    bool   `json:"available"`
}

type DishNameReq struct {
	RestaurantID int    `json:"restaurant_id"`
	OldName      string `json:"old_name"`
	NewName      string `json:"new_name"`
}

func getDishes(id int) []model.Dish {
	profiles := model.GetRestaurantProfilesFromFile()
	for i := 0; i < len(profiles.Profiles); i++ {
		if profiles.Profiles[i].ID == id {
			return profiles.Profiles[i].Dishes
		}
	}
	return nil
}

func updateDishes(id int, d []model.Dish) {
	profiles := model.GetRestaurantProfilesFromFile()
	for i := 0; i < len(profiles.Profiles); i++ {
		if profiles.Profiles[i].ID == id {
			profiles.Profiles[i].Dishes = d
			break
		}
	}
	model.UpdateRestaurantProfileFile(profiles)
}

func AddDish(c echo.Context) error {
	var dish DishReq
	if err := c.Bind(&dish); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	dishes := getDishes(dish.RestaurantID)
	for i := 0; i < len(dishes); i++ {
		if dishes[i].Name == dish.Name {
			return c.JSON(http.StatusConflict, ResponseMessage{
				StatusCode: http.StatusConflict,
				Message:    "A dish with this name already exists.",
			})
		}
	}

	newDish := model.Dish{
		Name:      dish.Name,
		Price:     dish.Price,
		Available: dish.Available,
	}

	profiles := model.GetRestaurantProfilesFromFile()
	for i := 0; i < len(profiles.Profiles); i++ {
		if profiles.Profiles[i].ID == dish.RestaurantID {
			profiles.Profiles[i].Dishes = append(profiles.Profiles[i].Dishes, newDish)
			break
		}
	}
	model.UpdateRestaurantProfileFile(profiles)
	return c.JSON(http.StatusOK, newDish)
}

func DeleteDish(c echo.Context) error {
	var dish DishReq
	if err := c.Bind(&dish); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	dishes := getDishes(dish.RestaurantID)
	for i := 0; i < len(dishes); i++ {
		if dishes[i].Name == dish.Name {
			dishes = append(dishes[:i], dishes[i+1:]...)
			break
		}
		if i == len(dishes)-1 {
			return c.JSON(http.StatusConflict, ResponseMessage{
				StatusCode: http.StatusBadRequest,
				Message:    "A dish with this name doesn't exist.",
			})
		}
	}

	updateDishes(dish.RestaurantID, dishes)

	return c.JSON(http.StatusOK, nil)
}

func UpdateDishPA(c echo.Context) error { //update dish price or availability
	var dish DishReq
	if err := c.Bind(&dish); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var updatedDish model.Dish
	dishes := getDishes(dish.RestaurantID)
	for i := 0; i < len(dishes); i++ {
		if dishes[i].Name == dish.Name {
			dishes[i].Price = dish.Price
			dishes[i].Available = dish.Available
			updatedDish = dishes[i]
			break
		}
		if i == len(dishes)-1 {
			return c.JSON(http.StatusConflict, ResponseMessage{
				StatusCode: http.StatusBadRequest,
				Message:    "A dish with this name doesn't exist.",
			})
		}
	}

	updateDishes(dish.RestaurantID, dishes)

	return c.JSON(http.StatusOK, updatedDish)
}

func UpdateDishName(c echo.Context) error {
	fmt.Println("hii")
	var dish DishNameReq
	if err := c.Bind(&dish); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	fmt.Println("hi")
	fmt.Println(dish.NewName)

	var updatedDish model.Dish
	dishes := getDishes(dish.RestaurantID)
	for i := 0; i < len(dishes); i++ {
		if dishes[i].Name == dish.NewName {
			return c.JSON(http.StatusConflict, ResponseMessage{
				StatusCode: http.StatusConflict,
				Message:    "A dish with this name already exists.",
			})
		}
		if dishes[i].Name == dish.OldName {
			dishes[i].Name = dish.NewName
			updatedDish = dishes[i]
			break
		}
		if i == len(dishes)-1 {
			return c.JSON(http.StatusConflict, ResponseMessage{
				StatusCode: http.StatusConflict,
				Message:    "A dish with this name doesn't exist.",
			})
		}
	}

	updateDishes(dish.RestaurantID, dishes)

	return c.JSON(http.StatusOK, updatedDish)
}
