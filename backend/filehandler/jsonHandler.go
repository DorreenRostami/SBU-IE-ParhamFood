package filehandler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Profiles struct {
	Profiles []Profile `json:"profiles"`
}

type Profile struct {
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

type Order struct {
	Customer  string `json:"customer"`
	Dishes    []Dish `json:"dishes"`
	Price     int    `json:"price"`
	Confirmed bool   `json:"confirmed"`
}

type Review struct {
	Customer string `json:"customer"`
	Text     string `json:"text"`
	Stars    int    `json:"stars"`
	Reply    string `json:"reply"`
}

func GetProfilesFromFile() Profiles {
	jsonFile, err := os.Open("resources/profiles.json")
	if err != nil {
		log.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	defer jsonFile.Close()
	var profiles Profiles
	json.Unmarshal(byteValue, &profiles)
	return profiles
}

func UpdateProfileFile(profiles Profiles) {
	file, _ := json.MarshalIndent(profiles, "", " ")
	_ = ioutil.WriteFile("resources/profiles.json", file, 0644)
}
