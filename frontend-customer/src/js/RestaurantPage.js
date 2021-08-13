import React, { useState, useEffect } from "react";
import "../css/Profile.css";
import axios from "axios";
import DishCard from "./DishCard";
import Button from "react-bootstrap/Button";

import { Link, useLocation } from "react-router-dom";

export default function RestaurantPage() {
  const [id, setId] = useState(0);
  const [resId, setResId] = useState(0);
  const location = useLocation();

  const [name, setName] = useState("");
  const [district, setDistrict] = useState("");
  const [address, setAddress] = useState("");
  const [open, setOpen] = useState(0);
  const [close, setClose] = useState(0);
  const [dish, setDish] = useState([]);
  const [fixedCost, setFixedCost] = useState(0);
  const [fixedMinute, setFixedMinute] = useState(0);

  useEffect(() => {
    setId(location.state.customer_id);
    setResId(location.state.res_id);
    const info = { restaurant_id: location.state.res_id };
    axios
      .post("http://localhost:1323/restaurantmenu", info)
      .then((response) => {
        setName(response.data.name);
        setDistrict(response.data.district);
        setAddress(response.data.address);
        setOpen(response.data.open);
        setClose(response.data.close);
        if (response.data.dishes == null) {
        } else setDish(response.data.dishes);
        console.log(response.data.dishes);
        setFixedCost(response.data.fixed_cost);
        setFixedMinute(response.data.fixed_minute);
      });
  }, []);

  var dishes = [];
  var i;
  var dishname;
  var price;
  var available;
  for (i = 0; i < dish.length; i++) {
    dishname = dish[i].name;
    price = dish[i].price;
    available = dish[i].available;

    dishes.push(
      <DishCard
        id={resId}
        name={dishname}
        price={price}
        available={available}
      />
    );
  }

  function order() {}

  return (
    <div className="Profile">
      <div
        style={{
          backgroundColor: "#ADD8E6",
          width: "20rem",
          margin: "2rem",
        }}
      >
        <Link
          to={{
            pathname: "/profile",
            state: {
              customer_id: id,
            },
          }}
        >
          Back to Profile
        </Link>
      </div>
      <div style={{ backgroundColor: "#D3D3D3" }}>
        <div>name: {name}</div>
        <div>district: {district}</div>
        <div>address: {address}</div>
        <div>open time: {open}</div>
        <div>close time: {close}</div>
        <div>fixed cost: {fixedCost}</div>
        <div>fixed minute: {fixedMinute}</div>
      </div>

      <div>{dishes}</div>
      <Button variant="primary" style={{ margin: "10px" }} onClick={order}>
        Order
      </Button>
    </div>
  );
}
