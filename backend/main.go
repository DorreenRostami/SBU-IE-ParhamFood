package main

import (
	rst "github.com/DorreenRostami/IE_ParhamFood/restaurant"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// var upgrader = websocket.Upgrader{
// 	CheckOrigin: func(r *http.Request) bool {
// 		return true
// 	},
// }

func main() {
	/*var data fh.RestaurantProfiles
	for i := 0; i < 2; i++ {
		v := strconv.Itoa(i)
		data.Profiles = append(data.Profiles, fh.RestaurantProfile{
			Email:    "admin" + v + "@gmail.com",
			Password: "1234",
			ID:       i,
			Name:     "sib" + v,
			District: "blaa",
			Address:  "bla bla bla",
			Open:     11,
			Close:    23,
			Dishes: []fh.Dish{
				{
					Name:      "ham",
					Price:     10,
					Available: true,
				},
			},
			FixedCost:   0,
			FixedMinute: 0,
			Orders: []fh.Order{
				{
					OrderID:      0,
					CustomerID:   0,
					RestaurantID: i,
					DisheInfos: []fh.DishInfo{
						{
							Name:     "ham",
							Price:    10,
							Quantity: 1,
						},
					},
					Price:     10,
					Status: 0
				},
			},
			Reviews: []fh.Review{},
		})
	}
	file, _ := json.MarshalIndent(data, "", " ")
	_ = ioutil.WriteFile("resources/profiles.json", file, 0644)*/

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
	e.POST("/loginadmin", rst.LoginAdmin)
	e.POST("/signupadmin", rst.SignUpAdmin)

	e.POST("/getmenu", rst.GetMenu)
	e.POST("/getorders", rst.GetOrders)
	e.POST("/getreviews", rst.GetReviews)
	e.POST("/getinfo", rst.GetInfo)

	e.POST("/adddish", rst.AddDish)
	e.POST("/deletedish", rst.DeleteDish)
	e.POST("/updatedishpa", rst.UpdateDishPA)
	e.POST("/updatedishname", rst.UpdateDishName)

	e.POST("/changeorderstatus", rst.ChangeOrderStatus)

	e.Logger.Fatal(e.Start(":1323"))

}

/*
 	data := Employee{
        FirstName: "Mark",
        LastName:  "Jones",
        Email:     "mark@gmail.com",
        Age:       25,
        MonthlySalary: []Salary{
            Salary{
                Basic: 15000.00,
                HRA:   5000.00,
                TA:    2000.00,
            },
            Salary{
                Basic: 16000.00,
                HRA:   5000.00,
                TA:    2100.00,
            },
            Salary{
                Basic: 17000.00,
                HRA:   5000.00,
                TA:    2200.00,
            },
        },
    }
*/
