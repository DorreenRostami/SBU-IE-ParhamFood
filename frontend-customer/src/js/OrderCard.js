import React from "react";

import Card from "react-bootstrap/Card";

function DishInfoCard(props) {
  return (
    <div
      className="DishInfoCard"
      style={{
        textAlign: "center",
        display: "flex",
        justifycontent: "center",
        alignitems: "center",
        backgroundColor: "black",
        margin: "20px",
      }}
    >
      <Card
        style={{ width: "18rem", backgroundColor: "#FCF6F5FF", margin: "1px" }}
      >
        <Card.Body>
          <Card.Text>Name: {props.name}</Card.Text>

          <Card.Text>Price: {props.price}</Card.Text>

          <Card.Text>Quantity: {props.quantity}</Card.Text>
        </Card.Body>
      </Card>
    </div>
  );
}

export default function OrderCard(props) {
  var infos = [];
  infos = [];
  var i;
  var name;
  var price;
  var quantity;
  for (i = 0; i < props.dishInfos.length; i++) {
    name = props.dishInfos[i].name;
    price = props.dishInfos[i].price;
    quantity = props.dishInfos[i].quantity;
    infos.push(<DishInfoCard name={name} price={price} quantity={quantity} />);
  }

  return (
    <div
      className="OrderCard"
      style={{
        width: "20rem",
        textAlign: "center",
        display: "flex",
        justifycontent: "center",
        alignitems: "center",
        backgroundColor: "#EA738DFF",
        margin: "20px",
      }}
    >
      <Card
        style={{ width: "18rem", backgroundColor: "#FCF6F5FF", margin: "15px" }}
      >
        <Card.Body>
          <Card.Title>Order ID: {props.orderID}</Card.Title>
          <Card.Text>Restaurant ID: {props.restaurantID}</Card.Text>
          <Card.Text>Dishes Information:</Card.Text>
          <div>{infos}</div>
          <Card.Text>Total Price: {props.price}</Card.Text>
          <Card.Text>Status: {props.status}</Card.Text>
          <Card.Text>Time: {props.time}</Card.Text>
        </Card.Body>
      </Card>
    </div>
  );
}
