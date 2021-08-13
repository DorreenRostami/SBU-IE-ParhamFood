import React, { useEffect, useState } from "react";
import OrderCard from "./OrderCard";

import axios from "axios";
import { Link, useLocation } from "react-router-dom";

export default function Orders() {
  const [id, setId] = useState(0);
  const [order, setOrder] = useState([]);

  const location = useLocation();
  useEffect(() => {
    setId(location.state.restaurant_id);
    const info = { restaurant_id: location.state.restaurant_id };
    axios.post("http://localhost:1323/getorders", info).then((response) => {
      setOrder(response.data.orders);
    });
  }, []);

  var orders = [];
  var i;
  var orderID;
  var customerID;
  var dishInfos;
  var totalPrice;
  var status;
  for (i = 0; i < order.length; i++) {
    orderID = order[i].order_id;
    customerID = order[i].customer_id;
    dishInfos = order[i].dishes;
    totalPrice = order[i].price;
    status = order[i].status;
    orders.push(
      <OrderCard
        orderID={orderID}
        customerID={customerID}
        dishInfos={dishInfos}
        totalPrice={totalPrice}
        status={status}
        id={location.state.restaurant_id}
      />
    );
  }

  return (
    <div className="Orders">
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
      {orders}
    </div>
  );
}
