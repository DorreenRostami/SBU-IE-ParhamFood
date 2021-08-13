import React, { useState, useEffect } from "react";
import "../css/Profile.css";
import axios from "axios";
import RestaurantCard from "./RestaurantCard";
import Button from "react-bootstrap/Button";

import { Link, useLocation } from "react-router-dom";

export default function Profile() {
  const [id, setId] = useState(0);
  const location = useLocation();

  const [searchText, setSearchText] = useState("");
  const [resultText, setResultText] = useState("List of All Restaurants:");
  const [restaurant, setRestaurant] = useState([]);

  useEffect(() => {
    setId(location.state.customer_id);
    axios.get("http://localhost:1323/homepage").then((response) => {
      setRestaurant(response.data.restaurant_homepage_infos);
    });
  }, []);

  var restaurants = [];
  var i;
  var restaurant_id;
  var name;
  var district;
  var address;
  for (i = 0; i < restaurant.length; i++) {
    restaurant_id = restaurant[i].restaurant_id;
    name = restaurant[i].name;
    district = restaurant[i].district;
    address = restaurant[i].address;
    restaurants.push(
      <RestaurantCard
        restaurant_id={restaurant_id}
        customer_id={id}
        name={name}
        district={district}
        address={address}
      />
    );
  }

  function foodSearch() {
    const info = {
      id: id,
      search_text: searchText,
    };
    axios.post("http://localhost:1323/searchbyfood", info).then((response) => {
      if (response.data.restaurant_homepage_infos == null) {
        setRestaurant([]);
      } else {
        setRestaurant(response.data.restaurant_homepage_infos);
      }
    });
    setResultText("Result by Food Name:");

    if (restaurant.length == 0) {
      restaurants = "Sorry, nothing found";
    } else {
      restaurants = [];
      var i;
      var restaurant_id;
      var name;
      var district;
      var address;
      for (i = 0; i < restaurant.length; i++) {
        restaurant_id = restaurant[i].restaurant_id;
        name = restaurant[i].name;
        district = restaurant[i].district;
        address = restaurant[i].address;
        restaurants.push(
          <RestaurantCard
            restaurant_id={restaurant_id}
            customer_id={id}
            name={name}
            district={district}
            address={address}
          />
        );
      }
    }
  }

  function restaurantSearch() {
    const info = {
      id: id,
      search_text: searchText,
    };
    axios.post("http://localhost:1323/searchbyname", info).then((response) => {
      if (response.data.restaurant_homepage_infos == null) {
        setRestaurant([]);
      } else {
        setRestaurant(response.data.restaurant_homepage_infos);
      }
    });
    setResultText("Result by Restaurant Name:");

    if (restaurant.length == 0) {
      restaurants = "Sorry, nothing found";
    } else {
      restaurants = [];
      var i;
      var restaurant_id;
      var name;
      var district;
      var address;
      for (i = 0; i < restaurant.length; i++) {
        restaurant_id = restaurant[i].restaurant_id;
        name = restaurant[i].name;
        district = restaurant[i].district;
        address = restaurant[i].address;
        restaurants.push(
          <RestaurantCard
            restaurant_id={restaurant_id}
            customer_id={id}
            name={name}
            district={district}
            address={address}
          />
        );
      }
    }
  }

  function districtSearch() {
    const info = {
      id: id,
      search_text: searchText,
    };
    axios
      .post("http://localhost:1323/searchbydistrict", info)
      .then((response) => {
        if (response.data.restaurant_homepage_infos == null) {
          setRestaurant([]);
        } else {
          setRestaurant(response.data.restaurant_homepage_infos);
        }
      });
    setResultText("Result by District:");
    if (restaurant.length == 0) {
      restaurants = "Sorry, nothing found";
    } else {
      restaurants = [];
      var i;
      var restaurant_id;
      var name;
      var district;
      var address;
      for (i = 0; i < restaurant.length; i++) {
        restaurant_id = restaurant[i].restaurant_id;
        name = restaurant[i].name;
        district = restaurant[i].district;
        address = restaurant[i].address;
        restaurants.push(
          <RestaurantCard
            restaurant_id={restaurant_id}
            customer_id={id}
            name={name}
            district={district}
            address={address}
          />
        );
      }
    }
  }

  function allSearch() {
    axios.get("http://localhost:1323/homepage").then((response) => {
      setRestaurant(response.data.restaurant_homepage_infos);
    });
    setResultText("List of All Restaurants:");
  }

  return (
    <div className="Profile" style={{ textAlign: "center" }}>
      <div
        style={{
          backgroundColor: "#ADD8E6",
          width: "20rem",
          margin: "2rem",
        }}
      >
        <Link
          to={{
            pathname: "/profile/edit",
            state: {
              customer_id: id,
            },
          }}
        >
          Get and Edit Your Information
        </Link>
      </div>
      <div
        style={{
          backgroundColor: "#ADD8E6",
          width: "20rem",
          margin: "2rem",
        }}
      >
        <Link
          to={{
            pathname: "/profile/ordershistory",
            state: {
              customer_id: id,
            },
          }}
        >
          Get the History of Your Orders
        </Link>
      </div>

      <div>
        <input
          type="text"
          name="reply"
          onChange={(e) => setSearchText(e.target.value)}
          style={{ margin: "10px", width: "20rem" }}
          placeholder="Search by Food, Restaurant or District..."
        />
        <Button
          variant="primary"
          style={{ margin: "10px" }}
          onClick={foodSearch}
        >
          Search by Food
        </Button>
        <Button
          variant="primary"
          style={{ margin: "10px" }}
          onClick={restaurantSearch}
        >
          Search by Restaurant
        </Button>

        <Button
          variant="primary"
          style={{ margin: "10px" }}
          onClick={districtSearch}
        >
          Search by District
        </Button>

        <Button
          variant="primary"
          style={{ margin: "10px" }}
          onClick={allSearch}
        >
          Get All Restaurants
        </Button>
      </div>

      <h2>{resultText}</h2>
      <div style={{ marginLeft: "37%" }}>{restaurants}</div>
    </div>
  );
}
