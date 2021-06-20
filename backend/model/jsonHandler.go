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

type CustomerProfiles struct {
	Profiles []CustomerProfile `json:"profiles"`
}

type CustomerProfile struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	District string `json:"district"`
	Address  string `json:"address"`
	Balance  int    `json:"balance"`
}

func GetRestaurantProfilesFromFile() RestaurantProfiles {
	jsonFile, err := os.Open("model/RestaurantProfiles.json")
	if err != nil {
		log.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	defer jsonFile.Close()
	var profiles RestaurantProfiles
	json.Unmarshal(byteValue, &profiles)
	return profiles
}

func UpdateRestaurantProfileFile(profiles RestaurantProfiles) {
	file, _ := json.MarshalIndent(profiles, "", " ")
	_ = ioutil.WriteFile("model/RestaurantProfiles.json", file, 0644)
}

func GetCustomerProfilesFromFile() CustomerProfiles {
	jsonFile, err := os.Open("model/CustomerProfiles.json")
	if err != nil {
		log.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	defer jsonFile.Close()
	var profiles CustomerProfiles
	json.Unmarshal(byteValue, &profiles)
	return profiles
}

func UpdateCustomerProfileFile(profiles CustomerProfiles) {
	file, _ := json.MarshalIndent(profiles, "", " ")
	_ = ioutil.WriteFile("model/CustomerProfiles.json", file, 0644)
}
