package model

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type RestaurantProfiles struct {
	Profiles []RestaurantProfile `json:"profiles"`
}

type RestaurantProfile struct {
	Email       string   `json:"email"`
	Password    string   `json:"password"`
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	District    string   `json:"district"`
	Address     string   `json:"address"`
	Open        int      `json:"open"`
	Close       int      `json:"close"`
	Dishes      []Dish   `json:"dishes"`
	FixedCost   int      `json:"fixed_cost"`
	FixedMinute int      `json:"fixed_minute"`
	Orders      []Order  `json:"orders"`
	Reviews     []Review `json:"review"`
}

type Dish struct {
	Name      string `json:"name"`
	Price     int    `json:"price"`
	Available bool   `json:"available"`
}

type DishInfo struct {
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
}

type Order struct {
	OrderID      int        `json:"order_id"`
	CustomerID   int        `json:"customer_id"`
	RestaurantID int        `json:"restaurant_id"`
	DisheInfos   []DishInfo `json:"dishes"`
	Price        int        `json:"price"`
	Status       int        `json:"status"`
}

type Review struct {
	ReviewID     int    `json:"review_id"`
	CustomerID   int    `json:"customer_id"`
	RestaurantID int    `json:"restaurant_id"`
	Text         string `json:"text"`
	Stars        int    `json:"stars"`
	Reply        string `json:"reply"`
}

func GetProfilesFromFile() RestaurantProfiles {
	jsonFile, err := os.Open("model/profiles.json")
	if err != nil {
		log.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	defer jsonFile.Close()
	var profiles RestaurantProfiles
	json.Unmarshal(byteValue, &profiles)
	return profiles
}

func UpdateProfileFile(profiles RestaurantProfiles) {
	file, _ := json.MarshalIndent(profiles, "", " ")
	_ = ioutil.WriteFile("model/profiles.json", file, 0644)
}
