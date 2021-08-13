import React from "react";
import Button from "react-bootstrap/Button";
import Card from "react-bootstrap/Card";
import axios from "axios";

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
  function nextStatus() {
    if (props.status == 4) {
    } else {
      var info = { order_id: props.order_id };
      axios.post("http://localhost:1323/changeorderstatus", info);
      //props.action(); // to do
    }
  }

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
          <Card.Text>Customer ID: {props.customerID}</Card.Text>
          <Card.Text>Dishes Information:</Card.Text>
          <div>{infos}</div>

          <Card.Text>Total Price: {props.totalPrice}</Card.Text>

          <Card.Text>Status: {props.status}</Card.Text>

          <div>
            <Button
              variant="primary"
              style={{ marginBottom: "10px" }}
              onClick={nextStatus}
            >
              Next Status
            </Button>
          </div>
        </Card.Body>
      </Card>
    </div>
  );
}
