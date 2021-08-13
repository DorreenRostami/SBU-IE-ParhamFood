import React, { useState, useEffect } from "react";
import "../css/Profile.css";
import axios from "axios";
import OrderCard from "./OrderCard";

import { Link, useLocation } from "react-router-dom";

export default function OrdersHistory() {
  const [id, setId] = useState(0);
  const [resId, setResId] = useState(0);
  const location = useLocation();

  const [order, setOrder] = useState([]);

  useEffect(() => {
    setId(location.state.customer_id);
    setResId(location.state.res_id);
    const info = { customer_id: location.state.customer_id };
    axios
      .post("http://localhost:1323/getcustomerorders", info)
      .then((response) => {
        if (response.data.orders == null) {
        } else setOrder(response.data.orders);
      });
  }, []);

  var orders = [];
  var i;
  var orderID;
  var restaurantID;
  var dishesInfo;
  var price;
  var status;
  var time;

  for (i = 0; i < order.length; i++) {
    orderID = order[i].order_id;
    restaurantID = order[i].restaurant_id;
    dishesInfo = order[i].dishes;
    price = order[i].price;
    status = order[i].status;
    time = order[i].time_of_order;
    orders.push(
      <OrderCard
        orderID={orderID}
        restaurantID={restaurantID}
        dishInfos={dishesInfo}
        price={price}
        status={status}
        time={time}
      />
    );
  }

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

      <div>{orders}</div>
    </div>
  );
}
