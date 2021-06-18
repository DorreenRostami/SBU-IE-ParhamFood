package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	rst "github.com/DorreenRostami/IE_ParhamFood/restaurant"
)

// var upgrader = websocket.Upgrader{
// 	CheckOrigin: func(r *http.Request) bool {
// 		return true
// 	},
// }

func main() {
	/*var data Profiles
	for i := 0; i < 3; i++ {
		v := strconv.Itoa(i)
		data.Profiles = append(data.Profiles, Profile{
			Email:     "admin" + v + "@gmail.com",
			Password:  "1234",
			ID:        i,
			Name:      "sib" + v,
			District:  "3",
			Address:   "bla bla bla",
			Open:      11,
			Close:     23,
			Dishes:    []Dish{},
			FixedCost: 0,
			FixedTime: 0,
			Orders:    []Order{},
			Reviews:   []Review{},
		})
	}
	file, _ := json.MarshalIndent(data, "", " ")
	_ = ioutil.WriteFile("profiles.json", file, 0644)

	jsonFile, err := os.Open("profiles.json")
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
		AllowOrigins: []string{"https://localhost:1323"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	//e.Static("/", "../src")
	e.GET("/loginadmin", rst.LoginAdmin)
	e.GET("/signupadmin", rst.SignUpAdmin)
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

	file, _ := json.MarshalIndent(data, "", " ")

	_ = ioutil.WriteFile("test.json", file, 0644)

	--------------------------------------------------------------------------------

	type j struct {
    Cl []string `json:"cl"`
    Gr []string `json:"gr"`
    Cr []string `json:"cr"`
}
*/
