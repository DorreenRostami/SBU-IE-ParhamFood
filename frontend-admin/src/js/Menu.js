import React, { useEffect, useState } from "react";
import Dishcard from "./Dishcard";
import "../css/Menu.css";
import axios from "axios";
import { Link, useLocation } from "react-router-dom";
import Card from "react-bootstrap/Card";
import Button from "react-bootstrap/Button";

export default function Menu() {
  const [id, setId] = useState(0);
  const [dish, setDish] = useState([]);

  const [newDishName, setNewDishName] = useState("");
  const [newDishPrice, setNewDishPrice] = useState(0);
  const [newDishStatus, setNewDishStatus] = useState("Available");

  const location = useLocation();
  useEffect(() => {
    setId(location.state.restaurant_id);
    const info = { restaurant_id: location.state.restaurant_id };
    axios.post("http://localhost:1323/getmenu", info).then((response) => {
      setDish(response.data.dishes);
    });
  }, []);

  function handler() {
    const info = { restaurant_id: location.state.restaurant_id };
    axios.post("http://localhost:1323/getmenu", info).then((response) => {
      setDish(response.data.dishes);
    });
  }

  function addDish() {
    if (newDishStatus == "Available") {
      var info = {
        restaurant_id: location.state.restaurant_id,
        name: newDishName,
        price: parseInt(newDishPrice, 10),
        available: true,
      };
    } else {
      info = {
        restaurant_id: location.state.restaurant_id,
        name: newDishName,
        price: parseInt(newDishPrice, 10),
        available: false,
      };
    }
    axios.post("http://localhost:1323/adddish", info);
  }

  var dishes = [];
  var i;
  var name;
  var price;
  var availibility;
  for (i = 0; i < dish.length; i++) {
    name = dish[i].name;
    price = dish[i].price;
    if (dish[i].available) {
      availibility = "available";
    } else {
      availibility = "not available";
    }
    dishes.push(
      <Dishcard
        name={name}
        availibility={availibility}
        price={price}
        action={handler}
        id={id}
      />
    );
  }

  return (
    <div className="Menu">
      <div>
        <Link
          to={{
            pathname: "/profile",
            state: {
              restaurant_id: id,
            },
          }}
          style={{ margin: "4rem" }}
        >
          Back to Profile
        </Link>
      </div>

      <div
        className="Dishcard"
        style={{
          width: "20rem",
          textAlign: "center",
          display: "flex",
          justifycontent: "center",
          alignitems: "center",
          backgroundColor: "#2BAE66FF",
          margin: "20px",
        }}
      >
        <Card
          style={{
            width: "18rem",
            backgroundColor: "#FCF6F5FF",
            margin: "20px",
          }}
        >
          <Card.Body>
            <Card.Title>Add a new dish to menu</Card.Title>
            <input
              type="text"
              name="newDishName"
              onChange={(e) => setNewDishName(e.target.value)}
              style={{ margin: "10px" }}
              placeholder="Enter new dish name"
            />
            <input
              type="number"
              name="newDishPrice"
              onChange={(e) => setNewDishPrice(e.target.value)}
              style={{ margin: "10px" }}
              placeholder="Enter new dish price"
            />

            <div>
              <input
                type="radio"
                defaultChecked
                value="Available"
                name="status"
                onChange={(event) => setNewDishStatus(event.target.value)}
              />{" "}
              Available
              <input
                type="radio"
                value="Unavailable"
                name="status"
                onChange={(event) => setNewDishStatus(event.target.value)}
              />{" "}
              Unavailable
            </div>
            <Button
              variant="primary"
              style={{ margin: "10px" }}
              onClick={addDish}
            >
              Add
            </Button>
          </Card.Body>
        </Card>
      </div>
      {dishes}
    </div>
  );
}
