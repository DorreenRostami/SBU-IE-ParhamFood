package main

import (
	cont "github.com/DorreenRostami/IE_ParhamFood/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	/*var data model.RestaurantProfiles
	for i := 0; i < 2; i++ {
		v := strconv.Itoa(i)
		data.Profiles = append(data.Profiles, model.RestaurantProfile{
			Email:    "admin" + v + "@gmail.com",
			Password: "1234",
			ID:       i,
			Name:     "sib" + v,
			District: "blaa",
			Address:  "bla bla bla",
			Open:     11,
			Close:    23,
			Dishes: []model.Dish{
				{
					Name:      "ham",
					Price:     10,
					Available: true,
				},
			},
			FixedCost:   0,
			FixedMinute: 0,
			Orders: []model.Order{
				{
					OrderID:      0,
					CustomerID:   0,
					RestaurantID: i,
					DisheInfos:   []model.DishInfo{{Name: "ham", Price: 10, Quantity: 1}},
					Price:        10,
					Status:       1,
				},
			},
			Reviews: []model.Review{},
		})
	}
	file, _ := json.MarshalIndent(data, "", " ")
	_ = ioutil.WriteFile("model/profiles.json", file, 0644)*/

	/*jsonFile, err := os.Open("profiles.json")
	if err != nil {
		log.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	defer jsonFile.Close()
	var profiles Profiles
	json.Unmarshal(byteValue, &profiles)
	for i := 0; i < len(profiles.Profiles); i++ {
		fmt.Println("email: " + profiles.Profiles[i].Email)
		fmt.Println("password: " + profiles.Profiles[i].Password)
		fmt.Println("restaurant name: " + profiles.Profiles[i].Name)
	}*/

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	//e.Static("/", "../src")
	e.POST("/loginadmin", cont.LoginAdmin)
	e.POST("/signupadmin", cont.SignUpAdmin)

	e.POST("/getmenu", cont.GetMenu)
	e.POST("/getorders", cont.GetOrders)
	e.POST("/getreviews", cont.GetReviews)
	e.POST("/getinfo", cont.GetInfo)

	e.POST("/adddish", cont.AddDish)
	e.POST("/deletedish", cont.DeleteDish)
	e.POST("/updatedishpa", cont.UpdateDishPA)
	e.POST("/updatedishname", cont.UpdateDishName)

	e.POST("/changeorderstatus", cont.ChangeOrderStatus)
	e.POST("/postreply", cont.PostReply)
	e.POST("/updateinfo", cont.UpdateInfo)

	e.Logger.Fatal(e.Start(":1323"))

}
